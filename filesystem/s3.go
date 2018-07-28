package filesystem

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/xitongsys/guery/config"
)

type S3FileSystem struct {
}

func SplitS3URL(url string) (prefix, bucket, key string, err error) {
	if strings.HasPrefix(url, "s3://") {
		prefix = "s3://"
	} else if strings.HasPrefix(url, "s3a://") {
		prefix = "s3a://"
	} else {
		err = fmt.Errorf("Unsupported s3 type")
		return
	}
	ns := strings.SplitN(url[len(prefix):], "/", 2)
	bucket, key = ns[0], ns[1]
	return
}

func (self *S3FileSystem) Accept(fl *FileLocation) bool {
	matched, err := regexp.MatchString("^s3.*://", fl.Location)
	if matched && err == nil {
		return true
	}
	return false
}

func (self *S3FileSystem) Open(fl *FileLocation) (VirtualFile, error) {
	var S3Conf = &aws.Config{
		Region: aws.String(Config.Conf.Runtime.S3Region),
	}
	svc := s3.New(session.New(), S3Conf)
	_, bucket, key, err := SplitS3URL(fl.Location)
	if err != nil {
		return nil, err
	}
	para := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	resp, err := svc.GetObject(para)
	if err != nil {
		return nil, err
	}

	tmpName := fmt.Sprintf("%s/s3_%d", os.TempDir(), rand.Uint32())
	tmpFile, err := os.Create(tmpName)
	if err != nil {
		return nil, err
	}
	size, err := io.Copy(tmpFile, resp.Body)
	resp.Body.Close()

	tmpFile.Seek(0, 0)
	return &VirtualFileS3{tmpFile, tmpName, size}, err

}

func (self *S3FileSystem) List(fl *FileLocation) (fileLocations []*FileLocation, err error) {
	res := []*FileLocation{}
	var S3Conf = &aws.Config{
		Region: aws.String(Config.Conf.Runtime.S3Region),
	}
	svc := s3.New(session.New(), S3Conf)
	prefix, bucket, key, err := SplitS3URL(fl.Location)
	if err != nil {
		return nil, err
	}
	para := &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(key),
	}

	resp, err := svc.ListObjects(para)
	if err != nil {
		return nil, err
	}

	for _, key := range resp.Contents {
		loc := fmt.Sprintf("%v%v/%v", prefix, bucket, *key.Key)
		res = append(res, NewFileLocation(loc, UNKNOWNFILETYPE))
	}
	return res, nil
}

func (self *S3FileSystem) IsDir(fl *FileLocation) bool {
	return false
}

type VirtualFileS3 struct {
	*os.File
	FileName string
	FileSize int64
}

func (self *VirtualFileS3) Size() int64 {
	return self.FileSize
}

func (self *VirtualFileS3) Close() error {
	self.File.Close()
	return os.Remove(self.FileName)
}
