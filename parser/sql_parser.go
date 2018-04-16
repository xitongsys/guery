// Code generated from ./Sql.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Sql

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 215, 601,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 5, 5, 101, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 107, 10, 6, 3, 7, 3,
	7, 3, 7, 3, 7, 5, 7, 113, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 119, 10,
	8, 12, 8, 14, 8, 122, 11, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 136, 10, 10, 12, 10, 14, 10,
	139, 11, 10, 5, 10, 141, 10, 10, 3, 10, 3, 10, 5, 10, 145, 10, 10, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 153, 10, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 159, 10, 11, 3, 11, 7, 11, 162, 10, 11, 12, 11, 14, 11,
	165, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 172, 10, 12, 3,
	13, 3, 13, 5, 13, 176, 10, 13, 3, 13, 3, 13, 5, 13, 180, 10, 13, 3, 14,
	3, 14, 5, 14, 184, 10, 14, 3, 14, 3, 14, 3, 14, 7, 14, 189, 10, 14, 12,
	14, 14, 14, 192, 11, 14, 3, 14, 3, 14, 3, 14, 3, 14, 7, 14, 198, 10, 14,
	12, 14, 14, 14, 201, 11, 14, 5, 14, 203, 10, 14, 3, 14, 3, 14, 5, 14, 207,
	10, 14, 3, 14, 3, 14, 3, 14, 5, 14, 212, 10, 14, 3, 14, 3, 14, 5, 14, 216,
	10, 14, 3, 15, 5, 15, 219, 10, 15, 3, 15, 3, 15, 3, 15, 7, 15, 224, 10,
	15, 12, 15, 14, 15, 227, 11, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3,
	18, 5, 18, 235, 10, 18, 3, 18, 5, 18, 238, 10, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 245, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 7, 19, 256, 10, 19, 12, 19, 14, 19, 259, 11, 19,
	3, 20, 5, 20, 262, 10, 20, 3, 20, 3, 20, 5, 20, 266, 10, 20, 3, 20, 3,
	20, 5, 20, 270, 10, 20, 3, 20, 3, 20, 5, 20, 274, 10, 20, 5, 20, 276, 10,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 285, 10, 21,
	12, 21, 14, 21, 288, 11, 21, 3, 21, 3, 21, 5, 21, 292, 10, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 5, 23, 298, 10, 23, 3, 23, 5, 23, 301, 10, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 312, 10,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 320, 10, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 7, 26, 328, 10, 26, 12, 26, 14, 26,
	331, 11, 26, 3, 27, 3, 27, 5, 27, 335, 10, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 347, 10, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 355, 10, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 3, 28, 7, 28, 362, 10, 28, 12, 28, 14, 28, 365, 11, 28, 3, 28,
	3, 28, 3, 28, 5, 28, 370, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3,
	28, 5, 28, 378, 10, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 384, 10, 28,
	3, 28, 3, 28, 5, 28, 388, 10, 28, 3, 28, 3, 28, 3, 28, 5, 28, 393, 10,
	28, 3, 28, 3, 28, 3, 28, 5, 28, 398, 10, 28, 3, 29, 3, 29, 3, 29, 3, 29,
	5, 29, 404, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 7, 29, 415, 10, 29, 12, 29, 14, 29, 418, 11, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 429, 10, 30, 3,
	30, 3, 30, 3, 30, 7, 30, 434, 10, 30, 12, 30, 14, 30, 437, 11, 30, 5, 30,
	439, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 7, 30, 446, 10, 30, 12,
	30, 14, 30, 449, 11, 30, 5, 30, 451, 10, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 5, 30, 459, 10, 30, 3, 30, 3, 30, 3, 30, 7, 30, 464, 10,
	30, 12, 30, 14, 30, 467, 11, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3,
	35, 3, 35, 3, 35, 7, 35, 498, 10, 35, 12, 35, 14, 35, 501, 11, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 510, 10, 35, 12, 35, 14,
	35, 513, 11, 35, 3, 35, 3, 35, 5, 35, 517, 10, 35, 5, 35, 519, 10, 35,
	3, 35, 3, 35, 7, 35, 523, 10, 35, 12, 35, 14, 35, 526, 11, 35, 3, 36, 3,
	36, 5, 36, 530, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 536, 10, 37,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 7, 40, 556, 10, 40,
	12, 40, 14, 40, 559, 11, 40, 5, 40, 561, 10, 40, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 40, 7, 40, 568, 10, 40, 12, 40, 14, 40, 571, 11, 40, 5, 40, 573,
	10, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 5, 41, 581, 10, 41, 3,
	42, 3, 42, 3, 42, 7, 42, 586, 10, 42, 12, 42, 14, 42, 589, 11, 42, 3, 43,
	3, 43, 3, 43, 3, 43, 5, 43, 595, 10, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3,
	45, 2, 8, 20, 36, 50, 56, 58, 68, 46, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56,
	58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 2, 16,
	4, 2, 53, 53, 76, 76, 4, 2, 8, 8, 202, 202, 4, 2, 52, 52, 172, 172, 4,
	2, 15, 15, 44, 44, 4, 2, 60, 60, 87, 87, 4, 2, 8, 8, 46, 46, 4, 2, 17,
	17, 155, 155, 3, 2, 193, 194, 3, 2, 195, 197, 3, 2, 187, 192, 5, 2, 8,
	8, 12, 12, 151, 151, 4, 2, 58, 58, 166, 166, 3, 2, 202, 203, 45, 2, 7,
	8, 10, 10, 12, 13, 15, 17, 20, 21, 24, 30, 35, 35, 39, 41, 44, 44, 47,
	47, 53, 53, 56, 56, 59, 61, 63, 63, 66, 69, 73, 74, 76, 76, 78, 78, 80,
	80, 82, 82, 85, 85, 87, 88, 90, 90, 92, 92, 95, 98, 100, 104, 108, 109,
	111, 112, 115, 115, 117, 122, 124, 128, 130, 135, 137, 137, 139, 143, 145,
	155, 157, 159, 161, 165, 167, 168, 170, 171, 174, 174, 176, 176, 178, 179,
	183, 186, 2, 654, 2, 90, 3, 2, 2, 2, 4, 93, 3, 2, 2, 2, 6, 96, 3, 2, 2,
	2, 8, 100, 3, 2, 2, 2, 10, 102, 3, 2, 2, 2, 12, 108, 3, 2, 2, 2, 14, 114,
	3, 2, 2, 2, 16, 125, 3, 2, 2, 2, 18, 129, 3, 2, 2, 2, 20, 146, 3, 2, 2,
	2, 22, 171, 3, 2, 2, 2, 24, 173, 3, 2, 2, 2, 26, 181, 3, 2, 2, 2, 28, 218,
	3, 2, 2, 2, 30, 228, 3, 2, 2, 2, 32, 230, 3, 2, 2, 2, 34, 244, 3, 2, 2,
	2, 36, 246, 3, 2, 2, 2, 38, 275, 3, 2, 2, 2, 40, 291, 3, 2, 2, 2, 42, 293,
	3, 2, 2, 2, 44, 295, 3, 2, 2, 2, 46, 311, 3, 2, 2, 2, 48, 313, 3, 2, 2,
	2, 50, 319, 3, 2, 2, 2, 52, 332, 3, 2, 2, 2, 54, 397, 3, 2, 2, 2, 56, 403,
	3, 2, 2, 2, 58, 458, 3, 2, 2, 2, 60, 468, 3, 2, 2, 2, 62, 470, 3, 2, 2,
	2, 64, 472, 3, 2, 2, 2, 66, 474, 3, 2, 2, 2, 68, 518, 3, 2, 2, 2, 70, 529,
	3, 2, 2, 2, 72, 535, 3, 2, 2, 2, 74, 537, 3, 2, 2, 2, 76, 542, 3, 2, 2,
	2, 78, 548, 3, 2, 2, 2, 80, 580, 3, 2, 2, 2, 82, 582, 3, 2, 2, 2, 84, 594,
	3, 2, 2, 2, 86, 596, 3, 2, 2, 2, 88, 598, 3, 2, 2, 2, 90, 91, 5, 6, 4,
	2, 91, 92, 7, 2, 2, 3, 92, 3, 3, 2, 2, 2, 93, 94, 5, 48, 25, 2, 94, 95,
	7, 2, 2, 3, 95, 5, 3, 2, 2, 2, 96, 97, 5, 18, 10, 2, 97, 7, 3, 2, 2, 2,
	98, 101, 5, 10, 6, 2, 99, 101, 5, 12, 7, 2, 100, 98, 3, 2, 2, 2, 100, 99,
	3, 2, 2, 2, 101, 9, 3, 2, 2, 2, 102, 103, 5, 84, 43, 2, 103, 106, 5, 68,
	35, 2, 104, 105, 7, 28, 2, 2, 105, 107, 5, 60, 31, 2, 106, 104, 3, 2, 2,
	2, 106, 107, 3, 2, 2, 2, 107, 11, 3, 2, 2, 2, 108, 109, 7, 91, 2, 2, 109,
	112, 5, 82, 42, 2, 110, 111, 9, 2, 2, 2, 111, 113, 7, 125, 2, 2, 112, 110,
	3, 2, 2, 2, 112, 113, 3, 2, 2, 2, 113, 13, 3, 2, 2, 2, 114, 115, 7, 3,
	2, 2, 115, 120, 5, 16, 9, 2, 116, 117, 7, 4, 2, 2, 117, 119, 5, 16, 9,
	2, 118, 116, 3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120,
	121, 3, 2, 2, 2, 121, 123, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124,
	7, 5, 2, 2, 124, 15, 3, 2, 2, 2, 125, 126, 5, 84, 43, 2, 126, 127, 7, 187,
	2, 2, 127, 128, 5, 48, 25, 2, 128, 17, 3, 2, 2, 2, 129, 140, 5, 20, 11,
	2, 130, 131, 7, 114, 2, 2, 131, 132, 7, 19, 2, 2, 132, 137, 5, 24, 13,
	2, 133, 134, 7, 4, 2, 2, 134, 136, 5, 24, 13, 2, 135, 133, 3, 2, 2, 2,
	136, 139, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 137, 138, 3, 2, 2, 2, 138,
	141, 3, 2, 2, 2, 139, 137, 3, 2, 2, 2, 140, 130, 3, 2, 2, 2, 140, 141,
	3, 2, 2, 2, 141, 144, 3, 2, 2, 2, 142, 143, 7, 92, 2, 2, 143, 145, 9, 3,
	2, 2, 144, 142, 3, 2, 2, 2, 144, 145, 3, 2, 2, 2, 145, 19, 3, 2, 2, 2,
	146, 147, 8, 11, 1, 2, 147, 148, 5, 22, 12, 2, 148, 163, 3, 2, 2, 2, 149,
	150, 12, 4, 2, 2, 150, 152, 7, 81, 2, 2, 151, 153, 5, 32, 17, 2, 152, 151,
	3, 2, 2, 2, 152, 153, 3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 162, 5, 20,
	11, 5, 155, 156, 12, 3, 2, 2, 156, 158, 9, 4, 2, 2, 157, 159, 5, 32, 17,
	2, 158, 157, 3, 2, 2, 2, 158, 159, 3, 2, 2, 2, 159, 160, 3, 2, 2, 2, 160,
	162, 5, 20, 11, 4, 161, 149, 3, 2, 2, 2, 161, 155, 3, 2, 2, 2, 162, 165,
	3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 21, 3, 2,
	2, 2, 165, 163, 3, 2, 2, 2, 166, 172, 5, 26, 14, 2, 167, 168, 7, 3, 2,
	2, 168, 169, 5, 18, 10, 2, 169, 170, 7, 5, 2, 2, 170, 172, 3, 2, 2, 2,
	171, 166, 3, 2, 2, 2, 171, 167, 3, 2, 2, 2, 172, 23, 3, 2, 2, 2, 173, 175,
	5, 48, 25, 2, 174, 176, 9, 5, 2, 2, 175, 174, 3, 2, 2, 2, 175, 176, 3,
	2, 2, 2, 176, 179, 3, 2, 2, 2, 177, 178, 7, 109, 2, 2, 178, 180, 9, 6,
	2, 2, 179, 177, 3, 2, 2, 2, 179, 180, 3, 2, 2, 2, 180, 25, 3, 2, 2, 2,
	181, 183, 7, 144, 2, 2, 182, 184, 5, 32, 17, 2, 183, 182, 3, 2, 2, 2, 183,
	184, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 190, 5, 34, 18, 2, 186, 187,
	7, 4, 2, 2, 187, 189, 5, 34, 18, 2, 188, 186, 3, 2, 2, 2, 189, 192, 3,
	2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 202, 3, 2, 2,
	2, 192, 190, 3, 2, 2, 2, 193, 194, 7, 64, 2, 2, 194, 199, 5, 36, 19, 2,
	195, 196, 7, 4, 2, 2, 196, 198, 5, 36, 19, 2, 197, 195, 3, 2, 2, 2, 198,
	201, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 199, 200, 3, 2, 2, 2, 200, 203,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 202, 193, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 206, 3, 2, 2, 2, 204, 205, 7, 181, 2, 2, 205, 207, 5, 50, 26,
	2, 206, 204, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 211, 3, 2, 2, 2, 208,
	209, 7, 70, 2, 2, 209, 210, 7, 19, 2, 2, 210, 212, 5, 28, 15, 2, 211, 208,
	3, 2, 2, 2, 211, 212, 3, 2, 2, 2, 212, 215, 3, 2, 2, 2, 213, 214, 7, 72,
	2, 2, 214, 216, 5, 50, 26, 2, 215, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2,
	2, 216, 27, 3, 2, 2, 2, 217, 219, 5, 32, 17, 2, 218, 217, 3, 2, 2, 2, 218,
	219, 3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 225, 5, 30, 16, 2, 221, 222,
	7, 4, 2, 2, 222, 224, 5, 30, 16, 2, 223, 221, 3, 2, 2, 2, 224, 227, 3,
	2, 2, 2, 225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 29, 3, 2, 2,
	2, 227, 225, 3, 2, 2, 2, 228, 229, 5, 48, 25, 2, 229, 31, 3, 2, 2, 2, 230,
	231, 9, 7, 2, 2, 231, 33, 3, 2, 2, 2, 232, 237, 5, 48, 25, 2, 233, 235,
	7, 14, 2, 2, 234, 233, 3, 2, 2, 2, 234, 235, 3, 2, 2, 2, 235, 236, 3, 2,
	2, 2, 236, 238, 5, 84, 43, 2, 237, 234, 3, 2, 2, 2, 237, 238, 3, 2, 2,
	2, 238, 245, 3, 2, 2, 2, 239, 240, 5, 82, 42, 2, 240, 241, 7, 6, 2, 2,
	241, 242, 7, 195, 2, 2, 242, 245, 3, 2, 2, 2, 243, 245, 7, 195, 2, 2, 244,
	232, 3, 2, 2, 2, 244, 239, 3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 35, 3,
	2, 2, 2, 246, 247, 8, 19, 1, 2, 247, 248, 5, 44, 23, 2, 248, 257, 3, 2,
	2, 2, 249, 250, 12, 4, 2, 2, 250, 251, 5, 38, 20, 2, 251, 252, 7, 86, 2,
	2, 252, 253, 5, 36, 19, 2, 253, 254, 5, 40, 21, 2, 254, 256, 3, 2, 2, 2,
	255, 249, 3, 2, 2, 2, 256, 259, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 257,
	258, 3, 2, 2, 2, 258, 37, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260, 262, 7,
	77, 2, 2, 261, 260, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 276, 3, 2, 2,
	2, 263, 265, 7, 89, 2, 2, 264, 266, 7, 116, 2, 2, 265, 264, 3, 2, 2, 2,
	265, 266, 3, 2, 2, 2, 266, 276, 3, 2, 2, 2, 267, 269, 7, 136, 2, 2, 268,
	270, 7, 116, 2, 2, 269, 268, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 276,
	3, 2, 2, 2, 271, 273, 7, 65, 2, 2, 272, 274, 7, 116, 2, 2, 273, 272, 3,
	2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 276, 3, 2, 2, 2, 275, 261, 3, 2, 2,
	2, 275, 263, 3, 2, 2, 2, 275, 267, 3, 2, 2, 2, 275, 271, 3, 2, 2, 2, 276,
	39, 3, 2, 2, 2, 277, 278, 7, 110, 2, 2, 278, 292, 5, 50, 26, 2, 279, 280,
	7, 175, 2, 2, 280, 281, 7, 3, 2, 2, 281, 286, 5, 84, 43, 2, 282, 283, 7,
	4, 2, 2, 283, 285, 5, 84, 43, 2, 284, 282, 3, 2, 2, 2, 285, 288, 3, 2,
	2, 2, 286, 284, 3, 2, 2, 2, 286, 287, 3, 2, 2, 2, 287, 289, 3, 2, 2, 2,
	288, 286, 3, 2, 2, 2, 289, 290, 7, 5, 2, 2, 290, 292, 3, 2, 2, 2, 291,
	277, 3, 2, 2, 2, 291, 279, 3, 2, 2, 2, 292, 41, 3, 2, 2, 2, 293, 294, 9,
	8, 2, 2, 294, 43, 3, 2, 2, 2, 295, 300, 5, 46, 24, 2, 296, 298, 7, 14,
	2, 2, 297, 296, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2,
	299, 301, 5, 84, 43, 2, 300, 297, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301,
	45, 3, 2, 2, 2, 302, 312, 5, 82, 42, 2, 303, 304, 7, 3, 2, 2, 304, 305,
	5, 18, 10, 2, 305, 306, 7, 5, 2, 2, 306, 312, 3, 2, 2, 2, 307, 308, 7,
	3, 2, 2, 308, 309, 5, 36, 19, 2, 309, 310, 7, 5, 2, 2, 310, 312, 3, 2,
	2, 2, 311, 302, 3, 2, 2, 2, 311, 303, 3, 2, 2, 2, 311, 307, 3, 2, 2, 2,
	312, 47, 3, 2, 2, 2, 313, 314, 5, 50, 26, 2, 314, 49, 3, 2, 2, 2, 315,
	316, 8, 26, 1, 2, 316, 320, 5, 52, 27, 2, 317, 318, 7, 106, 2, 2, 318,
	320, 5, 50, 26, 5, 319, 315, 3, 2, 2, 2, 319, 317, 3, 2, 2, 2, 320, 329,
	3, 2, 2, 2, 321, 322, 12, 4, 2, 2, 322, 323, 7, 11, 2, 2, 323, 328, 5,
	50, 26, 5, 324, 325, 12, 3, 2, 2, 325, 326, 7, 113, 2, 2, 326, 328, 5,
	50, 26, 4, 327, 321, 3, 2, 2, 2, 327, 324, 3, 2, 2, 2, 328, 331, 3, 2,
	2, 2, 329, 327, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 51, 3, 2, 2, 2,
	331, 329, 3, 2, 2, 2, 332, 334, 5, 56, 29, 2, 333, 335, 5, 54, 28, 2, 334,
	333, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 53, 3, 2, 2, 2, 336, 337, 5,
	62, 32, 2, 337, 338, 5, 56, 29, 2, 338, 398, 3, 2, 2, 2, 339, 340, 5, 62,
	32, 2, 340, 341, 5, 64, 33, 2, 341, 342, 7, 3, 2, 2, 342, 343, 5, 18, 10,
	2, 343, 344, 7, 5, 2, 2, 344, 398, 3, 2, 2, 2, 345, 347, 7, 106, 2, 2,
	346, 345, 3, 2, 2, 2, 346, 347, 3, 2, 2, 2, 347, 348, 3, 2, 2, 2, 348,
	349, 7, 18, 2, 2, 349, 350, 5, 56, 29, 2, 350, 351, 7, 11, 2, 2, 351, 352,
	5, 56, 29, 2, 352, 398, 3, 2, 2, 2, 353, 355, 7, 106, 2, 2, 354, 353, 3,
	2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 356, 3, 2, 2, 2, 356, 357, 7, 75, 2,
	2, 357, 358, 7, 3, 2, 2, 358, 363, 5, 48, 25, 2, 359, 360, 7, 4, 2, 2,
	360, 362, 5, 48, 25, 2, 361, 359, 3, 2, 2, 2, 362, 365, 3, 2, 2, 2, 363,
	361, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 366, 3, 2, 2, 2, 365, 363,
	3, 2, 2, 2, 366, 367, 7, 5, 2, 2, 367, 398, 3, 2, 2, 2, 368, 370, 7, 106,
	2, 2, 369, 368, 3, 2, 2, 2, 369, 370, 3, 2, 2, 2, 370, 371, 3, 2, 2, 2,
	371, 372, 7, 75, 2, 2, 372, 373, 7, 3, 2, 2, 373, 374, 5, 18, 10, 2, 374,
	375, 7, 5, 2, 2, 375, 398, 3, 2, 2, 2, 376, 378, 7, 106, 2, 2, 377, 376,
	3, 2, 2, 2, 377, 378, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380, 7, 91,
	2, 2, 380, 383, 5, 56, 29, 2, 381, 382, 7, 51, 2, 2, 382, 384, 5, 56, 29,
	2, 383, 381, 3, 2, 2, 2, 383, 384, 3, 2, 2, 2, 384, 398, 3, 2, 2, 2, 385,
	387, 7, 84, 2, 2, 386, 388, 7, 106, 2, 2, 387, 386, 3, 2, 2, 2, 387, 388,
	3, 2, 2, 2, 388, 389, 3, 2, 2, 2, 389, 398, 7, 107, 2, 2, 390, 392, 7,
	84, 2, 2, 391, 393, 7, 106, 2, 2, 392, 391, 3, 2, 2, 2, 392, 393, 3, 2,
	2, 2, 393, 394, 3, 2, 2, 2, 394, 395, 7, 46, 2, 2, 395, 396, 7, 64, 2,
	2, 396, 398, 5, 56, 29, 2, 397, 336, 3, 2, 2, 2, 397, 339, 3, 2, 2, 2,
	397, 346, 3, 2, 2, 2, 397, 354, 3, 2, 2, 2, 397, 369, 3, 2, 2, 2, 397,
	377, 3, 2, 2, 2, 397, 385, 3, 2, 2, 2, 397, 390, 3, 2, 2, 2, 398, 55, 3,
	2, 2, 2, 399, 400, 8, 29, 1, 2, 400, 404, 5, 58, 30, 2, 401, 402, 9, 9,
	2, 2, 402, 404, 5, 56, 29, 6, 403, 399, 3, 2, 2, 2, 403, 401, 3, 2, 2,
	2, 404, 416, 3, 2, 2, 2, 405, 406, 12, 5, 2, 2, 406, 407, 9, 10, 2, 2,
	407, 415, 5, 56, 29, 6, 408, 409, 12, 4, 2, 2, 409, 410, 9, 9, 2, 2, 410,
	415, 5, 56, 29, 5, 411, 412, 12, 3, 2, 2, 412, 413, 7, 198, 2, 2, 413,
	415, 5, 56, 29, 4, 414, 405, 3, 2, 2, 2, 414, 408, 3, 2, 2, 2, 414, 411,
	3, 2, 2, 2, 415, 418, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2,
	2, 2, 417, 57, 3, 2, 2, 2, 418, 416, 3, 2, 2, 2, 419, 420, 8, 30, 1, 2,
	420, 459, 7, 107, 2, 2, 421, 459, 5, 86, 44, 2, 422, 459, 5, 66, 34, 2,
	423, 459, 5, 60, 31, 2, 424, 459, 5, 84, 43, 2, 425, 426, 5, 82, 42, 2,
	426, 438, 7, 3, 2, 2, 427, 429, 5, 32, 17, 2, 428, 427, 3, 2, 2, 2, 428,
	429, 3, 2, 2, 2, 429, 430, 3, 2, 2, 2, 430, 435, 5, 48, 25, 2, 431, 432,
	7, 4, 2, 2, 432, 434, 5, 48, 25, 2, 433, 431, 3, 2, 2, 2, 434, 437, 3,
	2, 2, 2, 435, 433, 3, 2, 2, 2, 435, 436, 3, 2, 2, 2, 436, 439, 3, 2, 2,
	2, 437, 435, 3, 2, 2, 2, 438, 428, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439,
	450, 3, 2, 2, 2, 440, 441, 7, 114, 2, 2, 441, 442, 7, 19, 2, 2, 442, 447,
	5, 24, 13, 2, 443, 444, 7, 4, 2, 2, 444, 446, 5, 24, 13, 2, 445, 443, 3,
	2, 2, 2, 446, 449, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2,
	2, 448, 451, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 450, 440, 3, 2, 2, 2, 450,
	451, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 453, 7, 5, 2, 2, 453, 459,
	3, 2, 2, 2, 454, 455, 7, 3, 2, 2, 455, 456, 5, 48, 25, 2, 456, 457, 7,
	5, 2, 2, 457, 459, 3, 2, 2, 2, 458, 419, 3, 2, 2, 2, 458, 421, 3, 2, 2,
	2, 458, 422, 3, 2, 2, 2, 458, 423, 3, 2, 2, 2, 458, 424, 3, 2, 2, 2, 458,
	425, 3, 2, 2, 2, 458, 454, 3, 2, 2, 2, 459, 465, 3, 2, 2, 2, 460, 461,
	12, 4, 2, 2, 461, 462, 7, 6, 2, 2, 462, 464, 5, 84, 43, 2, 463, 460, 3,
	2, 2, 2, 464, 467, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 465, 466, 3, 2, 2,
	2, 466, 59, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 468, 469, 7, 199, 2, 2, 469,
	61, 3, 2, 2, 2, 470, 471, 9, 11, 2, 2, 471, 63, 3, 2, 2, 2, 472, 473, 9,
	12, 2, 2, 473, 65, 3, 2, 2, 2, 474, 475, 9, 13, 2, 2, 475, 67, 3, 2, 2,
	2, 476, 477, 8, 35, 1, 2, 477, 478, 7, 13, 2, 2, 478, 479, 7, 189, 2, 2,
	479, 480, 5, 68, 35, 2, 480, 481, 7, 191, 2, 2, 481, 519, 3, 2, 2, 2, 482,
	483, 7, 96, 2, 2, 483, 484, 7, 189, 2, 2, 484, 485, 5, 68, 35, 2, 485,
	486, 7, 4, 2, 2, 486, 487, 5, 68, 35, 2, 487, 488, 7, 191, 2, 2, 488, 519,
	3, 2, 2, 2, 489, 490, 7, 139, 2, 2, 490, 491, 7, 3, 2, 2, 491, 492, 5,
	84, 43, 2, 492, 499, 5, 68, 35, 2, 493, 494, 7, 4, 2, 2, 494, 495, 5, 84,
	43, 2, 495, 496, 5, 68, 35, 2, 496, 498, 3, 2, 2, 2, 497, 493, 3, 2, 2,
	2, 498, 501, 3, 2, 2, 2, 499, 497, 3, 2, 2, 2, 499, 500, 3, 2, 2, 2, 500,
	502, 3, 2, 2, 2, 501, 499, 3, 2, 2, 2, 502, 503, 7, 5, 2, 2, 503, 519,
	3, 2, 2, 2, 504, 516, 5, 72, 37, 2, 505, 506, 7, 3, 2, 2, 506, 511, 5,
	70, 36, 2, 507, 508, 7, 4, 2, 2, 508, 510, 5, 70, 36, 2, 509, 507, 3, 2,
	2, 2, 510, 513, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2, 511, 512, 3, 2, 2, 2,
	512, 514, 3, 2, 2, 2, 513, 511, 3, 2, 2, 2, 514, 515, 7, 5, 2, 2, 515,
	517, 3, 2, 2, 2, 516, 505, 3, 2, 2, 2, 516, 517, 3, 2, 2, 2, 517, 519,
	3, 2, 2, 2, 518, 476, 3, 2, 2, 2, 518, 482, 3, 2, 2, 2, 518, 489, 3, 2,
	2, 2, 518, 504, 3, 2, 2, 2, 519, 524, 3, 2, 2, 2, 520, 521, 12, 7, 2, 2,
	521, 523, 7, 13, 2, 2, 522, 520, 3, 2, 2, 2, 523, 526, 3, 2, 2, 2, 524,
	522, 3, 2, 2, 2, 524, 525, 3, 2, 2, 2, 525, 69, 3, 2, 2, 2, 526, 524, 3,
	2, 2, 2, 527, 530, 7, 202, 2, 2, 528, 530, 5, 68, 35, 2, 529, 527, 3, 2,
	2, 2, 529, 528, 3, 2, 2, 2, 530, 71, 3, 2, 2, 2, 531, 536, 7, 208, 2, 2,
	532, 536, 7, 209, 2, 2, 533, 536, 7, 210, 2, 2, 534, 536, 5, 84, 43, 2,
	535, 531, 3, 2, 2, 2, 535, 532, 3, 2, 2, 2, 535, 533, 3, 2, 2, 2, 535,
	534, 3, 2, 2, 2, 536, 73, 3, 2, 2, 2, 537, 538, 7, 180, 2, 2, 538, 539,
	5, 48, 25, 2, 539, 540, 7, 160, 2, 2, 540, 541, 5, 48, 25, 2, 541, 75,
	3, 2, 2, 2, 542, 543, 7, 59, 2, 2, 543, 544, 7, 3, 2, 2, 544, 545, 7, 181,
	2, 2, 545, 546, 5, 50, 26, 2, 546, 547, 7, 5, 2, 2, 547, 77, 3, 2, 2, 2,
	548, 549, 7, 118, 2, 2, 549, 560, 7, 3, 2, 2, 550, 551, 7, 119, 2, 2, 551,
	552, 7, 19, 2, 2, 552, 557, 5, 48, 25, 2, 553, 554, 7, 4, 2, 2, 554, 556,
	5, 48, 25, 2, 555, 553, 3, 2, 2, 2, 556, 559, 3, 2, 2, 2, 557, 555, 3,
	2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 561, 3, 2, 2, 2, 559, 557, 3, 2, 2,
	2, 560, 550, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561, 572, 3, 2, 2, 2, 562,
	563, 7, 114, 2, 2, 563, 564, 7, 19, 2, 2, 564, 569, 5, 24, 13, 2, 565,
	566, 7, 4, 2, 2, 566, 568, 5, 24, 13, 2, 567, 565, 3, 2, 2, 2, 568, 571,
	3, 2, 2, 2, 569, 567, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2, 570, 573, 3, 2,
	2, 2, 571, 569, 3, 2, 2, 2, 572, 562, 3, 2, 2, 2, 572, 573, 3, 2, 2, 2,
	573, 574, 3, 2, 2, 2, 574, 575, 7, 5, 2, 2, 575, 79, 3, 2, 2, 2, 576, 581,
	7, 144, 2, 2, 577, 581, 7, 43, 2, 2, 578, 581, 7, 79, 2, 2, 579, 581, 5,
	84, 43, 2, 580, 576, 3, 2, 2, 2, 580, 577, 3, 2, 2, 2, 580, 578, 3, 2,
	2, 2, 580, 579, 3, 2, 2, 2, 581, 81, 3, 2, 2, 2, 582, 587, 5, 84, 43, 2,
	583, 584, 7, 6, 2, 2, 584, 586, 5, 84, 43, 2, 585, 583, 3, 2, 2, 2, 586,
	589, 3, 2, 2, 2, 587, 585, 3, 2, 2, 2, 587, 588, 3, 2, 2, 2, 588, 83, 3,
	2, 2, 2, 589, 587, 3, 2, 2, 2, 590, 595, 7, 204, 2, 2, 591, 595, 7, 206,
	2, 2, 592, 595, 5, 88, 45, 2, 593, 595, 7, 205, 2, 2, 594, 590, 3, 2, 2,
	2, 594, 591, 3, 2, 2, 2, 594, 592, 3, 2, 2, 2, 594, 593, 3, 2, 2, 2, 595,
	85, 3, 2, 2, 2, 596, 597, 9, 14, 2, 2, 597, 87, 3, 2, 2, 2, 598, 599, 9,
	15, 2, 2, 599, 89, 3, 2, 2, 2, 76, 100, 106, 112, 120, 137, 140, 144, 152,
	158, 161, 163, 171, 175, 179, 183, 190, 199, 202, 206, 211, 215, 218, 225,
	234, 237, 244, 257, 261, 265, 269, 273, 275, 286, 291, 297, 300, 311, 319,
	327, 329, 334, 346, 354, 363, 369, 377, 383, 387, 392, 397, 403, 414, 416,
	428, 435, 438, 447, 450, 458, 465, 499, 511, 516, 518, 524, 529, 535, 557,
	560, 569, 572, 580, 587, 594,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'", "'.'", "'ADD'", "'ALL'", "'ALTER'", "'ANALYZE'",
	"'AND'", "'ANY'", "'ARRAY'", "'AS'", "'ASC'", "'AT'", "'BERNOULLI'", "'BETWEEN'",
	"'BY'", "'CALL'", "'CASCADE'", "'CASE'", "'CAST'", "'CATALOGS'", "'COALESCE'",
	"'COLUMN'", "'COLUMNS'", "'COMMENT'", "'COMMIT'", "'COMMITTED'", "'CONSTRAINT'",
	"'CREATE'", "'CROSS'", "'CUBE'", "'CURRENT'", "'CURRENT_DATE'", "'CURRENT_TIME'",
	"'CURRENT_TIMESTAMP'", "'DATA'", "'DATE'", "'DAY'", "'DEALLOCATE'", "'DELETE'",
	"'DESC'", "'DESCRIBE'", "'DISTINCT'", "'DISTRIBUTED'", "'DROP'", "'ELSE'",
	"'END'", "'ESCAPE'", "'EXCEPT'", "'EXCLUDING'", "'EXECUTE'", "'EXISTS'",
	"'EXPLAIN'", "'EXTRACT'", "'FALSE'", "'FILTER'", "'FIRST'", "'FOLLOWING'",
	"'FOR'", "'FORMAT'", "'FROM'", "'FULL'", "'FUNCTIONS'", "'GRANT'", "'GRANTS'",
	"'GRAPHVIZ'", "'GROUP'", "'GROUPING'", "'HAVING'", "'HOUR'", "'IF'", "'IN'",
	"'INCLUDING'", "'INNER'", "'INPUT'", "'INSERT'", "'INTEGER'", "'INTERSECT'",
	"'INTERVAL'", "'INTO'", "'IS'", "'ISOLATION'", "'JOIN'", "'LAST'", "'LATERAL'",
	"'LEFT'", "'LEVEL'", "'LIKE'", "'LIMIT'", "'LOCALTIME'", "'LOCALTIMESTAMP'",
	"'LOGICAL'", "'MAP'", "'MINUTE'", "'MONTH'", "'NATURAL'", "'NFC'", "'NFD'",
	"'NFKC'", "'NFKD'", "'NO'", "'NORMALIZE'", "'NOT'", "'NULL'", "'NULLIF'",
	"'NULLS'", "'ON'", "'ONLY'", "'OPTION'", "'OR'", "'ORDER'", "'ORDINALITY'",
	"'OUTER'", "'OUTPUT'", "'OVER'", "'PARTITION'", "'PARTITIONS'", "'POSITION'",
	"'PRECEDING'", "'PREPARE'", "'PRIVILEGES'", "'PROPERTIES'", "'PUBLIC'",
	"'RANGE'", "'READ'", "'RECURSIVE'", "'RENAME'", "'REPEATABLE'", "'REPLACE'",
	"'RESET'", "'RESTRICT'", "'REVOKE'", "'RIGHT'", "'ROLLBACK'", "'ROLLUP'",
	"'ROW'", "'ROWS'", "'SCHEMA'", "'SCHEMAS'", "'SECOND'", "'SELECT'", "'SERIALIZABLE'",
	"'SESSION'", "'SET'", "'SETS'", "'SHOW'", "'SMALLINT'", "'SOME'", "'START'",
	"'STATS'", "'SUBSTRING'", "'SYSTEM'", "'TABLE'", "'TABLES'", "'TABLESAMPLE'",
	"'TEXT'", "'THEN'", "'TIME'", "'TIMESTAMP'", "'TINYINT'", "'TO'", "'TRANSACTION'",
	"'TRUE'", "'TRY_CAST'", "'TYPE'", "'UESCAPE'", "'UNBOUNDED'", "'UNCOMMITTED'",
	"'UNION'", "'UNNEST'", "'USE'", "'USING'", "'VALIDATE'", "'VALUES'", "'VERBOSE'",
	"'VIEW'", "'WHEN'", "'WHERE'", "'WITH'", "'WORK'", "'WRITE'", "'YEAR'",
	"'ZONE'", "'='", "", "'<'", "'<='", "'>'", "'>='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'||'",
}
var symbolicNames = []string{
	"", "", "", "", "", "ADD", "ALL", "ALTER", "ANALYZE", "AND", "ANY", "ARRAY",
	"AS", "ASC", "AT", "BERNOULLI", "BETWEEN", "BY", "CALL", "CASCADE", "CASE",
	"CAST", "CATALOGS", "COALESCE", "COLUMN", "COLUMNS", "COMMENT", "COMMIT",
	"COMMITTED", "CONSTRAINT", "CREATE", "CROSS", "CUBE", "CURRENT", "CURRENT_DATE",
	"CURRENT_TIME", "CURRENT_TIMESTAMP", "DATA", "DATE", "DAY", "DEALLOCATE",
	"DELETE", "DESC", "DESCRIBE", "DISTINCT", "DISTRIBUTED", "DROP", "ELSE",
	"END", "ESCAPE", "EXCEPT", "EXCLUDING", "EXECUTE", "EXISTS", "EXPLAIN",
	"EXTRACT", "FALSE", "FILTER", "FIRST", "FOLLOWING", "FOR", "FORMAT", "FROM",
	"FULL", "FUNCTIONS", "GRANT", "GRANTS", "GRAPHVIZ", "GROUP", "GROUPING",
	"HAVING", "HOUR", "IF", "IN", "INCLUDING", "INNER", "INPUT", "INSERT",
	"INTEGER", "INTERSECT", "INTERVAL", "INTO", "IS", "ISOLATION", "JOIN",
	"LAST", "LATERAL", "LEFT", "LEVEL", "LIKE", "LIMIT", "LOCALTIME", "LOCALTIMESTAMP",
	"LOGICAL", "MAP", "MINUTE", "MONTH", "NATURAL", "NFC", "NFD", "NFKC", "NFKD",
	"NO", "NORMALIZE", "NOT", "NULL", "NULLIF", "NULLS", "ON", "ONLY", "OPTION",
	"OR", "ORDER", "ORDINALITY", "OUTER", "OUTPUT", "OVER", "PARTITION", "PARTITIONS",
	"POSITION", "PRECEDING", "PREPARE", "PRIVILEGES", "PROPERTIES", "PUBLIC",
	"RANGE", "READ", "RECURSIVE", "RENAME", "REPEATABLE", "REPLACE", "RESET",
	"RESTRICT", "REVOKE", "RIGHT", "ROLLBACK", "ROLLUP", "ROW", "ROWS", "SCHEMA",
	"SCHEMAS", "SECOND", "SELECT", "SERIALIZABLE", "SESSION", "SET", "SETS",
	"SHOW", "SMALLINT", "SOME", "START", "STATS", "SUBSTRING", "SYSTEM", "TABLE",
	"TABLES", "TABLESAMPLE", "TEXT", "THEN", "TIME", "TIMESTAMP", "TINYINT",
	"TO", "TRANSACTION", "TRUE", "TRY_CAST", "TYPE", "UESCAPE", "UNBOUNDED",
	"UNCOMMITTED", "UNION", "UNNEST", "USE", "USING", "VALIDATE", "VALUES",
	"VERBOSE", "VIEW", "WHEN", "WHERE", "WITH", "WORK", "WRITE", "YEAR", "ZONE",
	"EQ", "NEQ", "LT", "LTE", "GT", "GTE", "PLUS", "MINUS", "ASTERISK", "SLASH",
	"PERCENT", "CONCAT", "STRING", "UNICODE_STRING", "BINARY_LITERAL", "INTEGER_VALUE",
	"DOUBLE_VALUE", "IDENTIFIER", "DIGIT_IDENTIFIER", "QUOTED_IDENTIFIER",
	"BACKQUOTED_IDENTIFIER", "TIME_WITH_TIME_ZONE", "TIMESTAMP_WITH_TIME_ZONE",
	"DOUBLE_PRECISION", "SIMPLE_COMMENT", "BRACKETED_COMMENT", "WS", "UNRECOGNIZED",
	"DELIMITER",
}

var ruleNames = []string{
	"singleStatement", "singleExpression", "statement", "tableElement", "columnDefinition",
	"likeClause", "properties", "property", "query", "queryTerm", "queryPrimary",
	"sortItem", "querySpecification", "groupBy", "groupingElement", "setQuantifier",
	"selectItem", "relation", "joinType", "joinCriteria", "sampleType", "sampledRelation",
	"relationPrimary", "expression", "booleanExpression", "predicated", "predicate",
	"valueExpression", "primaryExpression", "stringValue", "comparisonOperator",
	"comparisonQuantifier", "booleanValue", "typeSql", "typeParameter", "baseType",
	"whenClause", "filter", "over", "privilege", "qualifiedName", "identifier",
	"number", "nonReserved",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SqlParser struct {
	*antlr.BaseParser
}

func NewSqlParser(input antlr.TokenStream) *SqlParser {
	this := new(SqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sql.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF                      = antlr.TokenEOF
	SqlParserT__0                     = 1
	SqlParserT__1                     = 2
	SqlParserT__2                     = 3
	SqlParserT__3                     = 4
	SqlParserADD                      = 5
	SqlParserALL                      = 6
	SqlParserALTER                    = 7
	SqlParserANALYZE                  = 8
	SqlParserAND                      = 9
	SqlParserANY                      = 10
	SqlParserARRAY                    = 11
	SqlParserAS                       = 12
	SqlParserASC                      = 13
	SqlParserAT                       = 14
	SqlParserBERNOULLI                = 15
	SqlParserBETWEEN                  = 16
	SqlParserBY                       = 17
	SqlParserCALL                     = 18
	SqlParserCASCADE                  = 19
	SqlParserCASE                     = 20
	SqlParserCAST                     = 21
	SqlParserCATALOGS                 = 22
	SqlParserCOALESCE                 = 23
	SqlParserCOLUMN                   = 24
	SqlParserCOLUMNS                  = 25
	SqlParserCOMMENT                  = 26
	SqlParserCOMMIT                   = 27
	SqlParserCOMMITTED                = 28
	SqlParserCONSTRAINT               = 29
	SqlParserCREATE                   = 30
	SqlParserCROSS                    = 31
	SqlParserCUBE                     = 32
	SqlParserCURRENT                  = 33
	SqlParserCURRENT_DATE             = 34
	SqlParserCURRENT_TIME             = 35
	SqlParserCURRENT_TIMESTAMP        = 36
	SqlParserDATA                     = 37
	SqlParserDATE                     = 38
	SqlParserDAY                      = 39
	SqlParserDEALLOCATE               = 40
	SqlParserDELETE                   = 41
	SqlParserDESC                     = 42
	SqlParserDESCRIBE                 = 43
	SqlParserDISTINCT                 = 44
	SqlParserDISTRIBUTED              = 45
	SqlParserDROP                     = 46
	SqlParserELSE                     = 47
	SqlParserEND                      = 48
	SqlParserESCAPE                   = 49
	SqlParserEXCEPT                   = 50
	SqlParserEXCLUDING                = 51
	SqlParserEXECUTE                  = 52
	SqlParserEXISTS                   = 53
	SqlParserEXPLAIN                  = 54
	SqlParserEXTRACT                  = 55
	SqlParserFALSE                    = 56
	SqlParserFILTER                   = 57
	SqlParserFIRST                    = 58
	SqlParserFOLLOWING                = 59
	SqlParserFOR                      = 60
	SqlParserFORMAT                   = 61
	SqlParserFROM                     = 62
	SqlParserFULL                     = 63
	SqlParserFUNCTIONS                = 64
	SqlParserGRANT                    = 65
	SqlParserGRANTS                   = 66
	SqlParserGRAPHVIZ                 = 67
	SqlParserGROUP                    = 68
	SqlParserGROUPING                 = 69
	SqlParserHAVING                   = 70
	SqlParserHOUR                     = 71
	SqlParserIF                       = 72
	SqlParserIN                       = 73
	SqlParserINCLUDING                = 74
	SqlParserINNER                    = 75
	SqlParserINPUT                    = 76
	SqlParserINSERT                   = 77
	SqlParserINTEGER                  = 78
	SqlParserINTERSECT                = 79
	SqlParserINTERVAL                 = 80
	SqlParserINTO                     = 81
	SqlParserIS                       = 82
	SqlParserISOLATION                = 83
	SqlParserJOIN                     = 84
	SqlParserLAST                     = 85
	SqlParserLATERAL                  = 86
	SqlParserLEFT                     = 87
	SqlParserLEVEL                    = 88
	SqlParserLIKE                     = 89
	SqlParserLIMIT                    = 90
	SqlParserLOCALTIME                = 91
	SqlParserLOCALTIMESTAMP           = 92
	SqlParserLOGICAL                  = 93
	SqlParserMAP                      = 94
	SqlParserMINUTE                   = 95
	SqlParserMONTH                    = 96
	SqlParserNATURAL                  = 97
	SqlParserNFC                      = 98
	SqlParserNFD                      = 99
	SqlParserNFKC                     = 100
	SqlParserNFKD                     = 101
	SqlParserNO                       = 102
	SqlParserNORMALIZE                = 103
	SqlParserNOT                      = 104
	SqlParserNULL                     = 105
	SqlParserNULLIF                   = 106
	SqlParserNULLS                    = 107
	SqlParserON                       = 108
	SqlParserONLY                     = 109
	SqlParserOPTION                   = 110
	SqlParserOR                       = 111
	SqlParserORDER                    = 112
	SqlParserORDINALITY               = 113
	SqlParserOUTER                    = 114
	SqlParserOUTPUT                   = 115
	SqlParserOVER                     = 116
	SqlParserPARTITION                = 117
	SqlParserPARTITIONS               = 118
	SqlParserPOSITION                 = 119
	SqlParserPRECEDING                = 120
	SqlParserPREPARE                  = 121
	SqlParserPRIVILEGES               = 122
	SqlParserPROPERTIES               = 123
	SqlParserPUBLIC                   = 124
	SqlParserRANGE                    = 125
	SqlParserREAD                     = 126
	SqlParserRECURSIVE                = 127
	SqlParserRENAME                   = 128
	SqlParserREPEATABLE               = 129
	SqlParserREPLACE                  = 130
	SqlParserRESET                    = 131
	SqlParserRESTRICT                 = 132
	SqlParserREVOKE                   = 133
	SqlParserRIGHT                    = 134
	SqlParserROLLBACK                 = 135
	SqlParserROLLUP                   = 136
	SqlParserROW                      = 137
	SqlParserROWS                     = 138
	SqlParserSCHEMA                   = 139
	SqlParserSCHEMAS                  = 140
	SqlParserSECOND                   = 141
	SqlParserSELECT                   = 142
	SqlParserSERIALIZABLE             = 143
	SqlParserSESSION                  = 144
	SqlParserSET                      = 145
	SqlParserSETS                     = 146
	SqlParserSHOW                     = 147
	SqlParserSMALLINT                 = 148
	SqlParserSOME                     = 149
	SqlParserSTART                    = 150
	SqlParserSTATS                    = 151
	SqlParserSUBSTRING                = 152
	SqlParserSYSTEM                   = 153
	SqlParserTABLE                    = 154
	SqlParserTABLES                   = 155
	SqlParserTABLESAMPLE              = 156
	SqlParserTEXT                     = 157
	SqlParserTHEN                     = 158
	SqlParserTIME                     = 159
	SqlParserTIMESTAMP                = 160
	SqlParserTINYINT                  = 161
	SqlParserTO                       = 162
	SqlParserTRANSACTION              = 163
	SqlParserTRUE                     = 164
	SqlParserTRY_CAST                 = 165
	SqlParserTYPE                     = 166
	SqlParserUESCAPE                  = 167
	SqlParserUNBOUNDED                = 168
	SqlParserUNCOMMITTED              = 169
	SqlParserUNION                    = 170
	SqlParserUNNEST                   = 171
	SqlParserUSE                      = 172
	SqlParserUSING                    = 173
	SqlParserVALIDATE                 = 174
	SqlParserVALUES                   = 175
	SqlParserVERBOSE                  = 176
	SqlParserVIEW                     = 177
	SqlParserWHEN                     = 178
	SqlParserWHERE                    = 179
	SqlParserWITH                     = 180
	SqlParserWORK                     = 181
	SqlParserWRITE                    = 182
	SqlParserYEAR                     = 183
	SqlParserZONE                     = 184
	SqlParserEQ                       = 185
	SqlParserNEQ                      = 186
	SqlParserLT                       = 187
	SqlParserLTE                      = 188
	SqlParserGT                       = 189
	SqlParserGTE                      = 190
	SqlParserPLUS                     = 191
	SqlParserMINUS                    = 192
	SqlParserASTERISK                 = 193
	SqlParserSLASH                    = 194
	SqlParserPERCENT                  = 195
	SqlParserCONCAT                   = 196
	SqlParserSTRING                   = 197
	SqlParserUNICODE_STRING           = 198
	SqlParserBINARY_LITERAL           = 199
	SqlParserINTEGER_VALUE            = 200
	SqlParserDOUBLE_VALUE             = 201
	SqlParserIDENTIFIER               = 202
	SqlParserDIGIT_IDENTIFIER         = 203
	SqlParserQUOTED_IDENTIFIER        = 204
	SqlParserBACKQUOTED_IDENTIFIER    = 205
	SqlParserTIME_WITH_TIME_ZONE      = 206
	SqlParserTIMESTAMP_WITH_TIME_ZONE = 207
	SqlParserDOUBLE_PRECISION         = 208
	SqlParserSIMPLE_COMMENT           = 209
	SqlParserBRACKETED_COMMENT        = 210
	SqlParserWS                       = 211
	SqlParserUNRECOGNIZED             = 212
	SqlParserDELIMITER                = 213
)

// SqlParser rules.
const (
	SqlParserRULE_singleStatement      = 0
	SqlParserRULE_singleExpression     = 1
	SqlParserRULE_statement            = 2
	SqlParserRULE_tableElement         = 3
	SqlParserRULE_columnDefinition     = 4
	SqlParserRULE_likeClause           = 5
	SqlParserRULE_properties           = 6
	SqlParserRULE_property             = 7
	SqlParserRULE_query                = 8
	SqlParserRULE_queryTerm            = 9
	SqlParserRULE_queryPrimary         = 10
	SqlParserRULE_sortItem             = 11
	SqlParserRULE_querySpecification   = 12
	SqlParserRULE_groupBy              = 13
	SqlParserRULE_groupingElement      = 14
	SqlParserRULE_setQuantifier        = 15
	SqlParserRULE_selectItem           = 16
	SqlParserRULE_relation             = 17
	SqlParserRULE_joinType             = 18
	SqlParserRULE_joinCriteria         = 19
	SqlParserRULE_sampleType           = 20
	SqlParserRULE_sampledRelation      = 21
	SqlParserRULE_relationPrimary      = 22
	SqlParserRULE_expression           = 23
	SqlParserRULE_booleanExpression    = 24
	SqlParserRULE_predicated           = 25
	SqlParserRULE_predicate            = 26
	SqlParserRULE_valueExpression      = 27
	SqlParserRULE_primaryExpression    = 28
	SqlParserRULE_stringValue          = 29
	SqlParserRULE_comparisonOperator   = 30
	SqlParserRULE_comparisonQuantifier = 31
	SqlParserRULE_booleanValue         = 32
	SqlParserRULE_typeSql              = 33
	SqlParserRULE_typeParameter        = 34
	SqlParserRULE_baseType             = 35
	SqlParserRULE_whenClause           = 36
	SqlParserRULE_filter               = 37
	SqlParserRULE_over                 = 38
	SqlParserRULE_privilege            = 39
	SqlParserRULE_qualifiedName        = 40
	SqlParserRULE_identifier           = 41
	SqlParserRULE_number               = 42
	SqlParserRULE_nonReserved          = 43
)

// ISingleStatementContext is an interface to support dynamic dispatch.
type ISingleStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleStatementContext differentiates from other interfaces.
	IsSingleStatementContext()
}

type SingleStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleStatementContext() *SingleStatementContext {
	var p = new(SingleStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_singleStatement
	return p
}

func (*SingleStatementContext) IsSingleStatementContext() {}

func NewSingleStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleStatementContext {
	var p = new(SingleStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_singleStatement

	return p
}

func (s *SingleStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SingleStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SqlParserEOF, 0)
}

func (s *SingleStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSingleStatement(s)
	}
}

func (s *SingleStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSingleStatement(s)
	}
}

func (p *SqlParser) SingleStatement() (localctx ISingleStatementContext) {
	localctx = NewSingleStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_singleStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(88)
		p.Statement()
	}
	{
		p.SetState(89)
		p.Match(SqlParserEOF)
	}

	return localctx
}

// ISingleExpressionContext is an interface to support dynamic dispatch.
type ISingleExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleExpressionContext differentiates from other interfaces.
	IsSingleExpressionContext()
}

type SingleExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleExpressionContext() *SingleExpressionContext {
	var p = new(SingleExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_singleExpression
	return p
}

func (*SingleExpressionContext) IsSingleExpressionContext() {}

func NewSingleExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleExpressionContext {
	var p = new(SingleExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_singleExpression

	return p
}

func (s *SingleExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SingleExpressionContext) EOF() antlr.TerminalNode {
	return s.GetToken(SqlParserEOF, 0)
}

func (s *SingleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSingleExpression(s)
	}
}

func (s *SingleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSingleExpression(s)
	}
}

func (p *SqlParser) SingleExpression() (localctx ISingleExpressionContext) {
	localctx = NewSingleExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_singleExpression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Expression()
	}
	{
		p.SetState(92)
		p.Match(SqlParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *SqlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(94)
		p.Query()
	}

	return localctx
}

// ITableElementContext is an interface to support dynamic dispatch.
type ITableElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableElementContext differentiates from other interfaces.
	IsTableElementContext()
}

type TableElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableElementContext() *TableElementContext {
	var p = new(TableElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableElement
	return p
}

func (*TableElementContext) IsTableElementContext() {}

func NewTableElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableElementContext {
	var p = new(TableElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableElement

	return p
}

func (s *TableElementContext) GetParser() antlr.Parser { return s.parser }

func (s *TableElementContext) ColumnDefinition() IColumnDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnDefinitionContext)
}

func (s *TableElementContext) LikeClause() ILikeClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILikeClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILikeClauseContext)
}

func (s *TableElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTableElement(s)
	}
}

func (s *TableElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTableElement(s)
	}
}

func (p *SqlParser) TableElement() (localctx ITableElementContext) {
	localctx = NewTableElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_tableElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(96)
			p.ColumnDefinition()
		}

	case SqlParserLIKE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(97)
			p.LikeClause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IColumnDefinitionContext is an interface to support dynamic dispatch.
type IColumnDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnDefinitionContext differentiates from other interfaces.
	IsColumnDefinitionContext()
}

type ColumnDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnDefinitionContext() *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnDefinition
	return p
}

func (*ColumnDefinitionContext) IsColumnDefinitionContext() {}

func NewColumnDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnDefinitionContext {
	var p = new(ColumnDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnDefinition

	return p
}

func (s *ColumnDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnDefinitionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *ColumnDefinitionContext) TypeSql() ITypeSqlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSqlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSqlContext)
}

func (s *ColumnDefinitionContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMENT, 0)
}

func (s *ColumnDefinitionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ColumnDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterColumnDefinition(s)
	}
}

func (s *ColumnDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitColumnDefinition(s)
	}
}

func (p *SqlParser) ColumnDefinition() (localctx IColumnDefinitionContext) {
	localctx = NewColumnDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_columnDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(100)
		p.Identifier()
	}
	{
		p.SetState(101)
		p.typeSql(0)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserCOMMENT {
		{
			p.SetState(102)
			p.Match(SqlParserCOMMENT)
		}
		{
			p.SetState(103)
			p.StringValue()
		}

	}

	return localctx
}

// ILikeClauseContext is an interface to support dynamic dispatch.
type ILikeClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOptionType returns the optionType token.
	GetOptionType() antlr.Token

	// SetOptionType sets the optionType token.
	SetOptionType(antlr.Token)

	// IsLikeClauseContext differentiates from other interfaces.
	IsLikeClauseContext()
}

type LikeClauseContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	optionType antlr.Token
}

func NewEmptyLikeClauseContext() *LikeClauseContext {
	var p = new(LikeClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_likeClause
	return p
}

func (*LikeClauseContext) IsLikeClauseContext() {}

func NewLikeClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LikeClauseContext {
	var p = new(LikeClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_likeClause

	return p
}

func (s *LikeClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LikeClauseContext) GetOptionType() antlr.Token { return s.optionType }

func (s *LikeClauseContext) SetOptionType(v antlr.Token) { s.optionType = v }

func (s *LikeClauseContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserLIKE, 0)
}

func (s *LikeClauseContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *LikeClauseContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(SqlParserPROPERTIES, 0)
}

func (s *LikeClauseContext) INCLUDING() antlr.TerminalNode {
	return s.GetToken(SqlParserINCLUDING, 0)
}

func (s *LikeClauseContext) EXCLUDING() antlr.TerminalNode {
	return s.GetToken(SqlParserEXCLUDING, 0)
}

func (s *LikeClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikeClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LikeClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterLikeClause(s)
	}
}

func (s *LikeClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitLikeClause(s)
	}
}

func (p *SqlParser) LikeClause() (localctx ILikeClauseContext) {
	localctx = NewLikeClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_likeClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Match(SqlParserLIKE)
	}
	{
		p.SetState(107)
		p.QualifiedName()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserEXCLUDING || _la == SqlParserINCLUDING {
		{
			p.SetState(108)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*LikeClauseContext).optionType = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserEXCLUDING || _la == SqlParserINCLUDING) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*LikeClauseContext).optionType = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(109)
			p.Match(SqlParserPROPERTIES)
		}

	}

	return localctx
}

// IPropertiesContext is an interface to support dynamic dispatch.
type IPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertiesContext differentiates from other interfaces.
	IsPropertiesContext()
}

type PropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertiesContext() *PropertiesContext {
	var p = new(PropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_properties
	return p
}

func (*PropertiesContext) IsPropertiesContext() {}

func NewPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertiesContext {
	var p = new(PropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_properties

	return p
}

func (s *PropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertiesContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *PropertiesContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertiesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterProperties(s)
	}
}

func (s *PropertiesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitProperties(s)
	}
}

func (p *SqlParser) Properties() (localctx IPropertiesContext) {
	localctx = NewPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_properties)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(SqlParserT__0)
	}
	{
		p.SetState(113)
		p.Property()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserT__1 {
		{
			p.SetState(114)
			p.Match(SqlParserT__1)
		}
		{
			p.SetState(115)
			p.Property()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(121)
		p.Match(SqlParserT__2)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PropertyContext) EQ() antlr.TerminalNode {
	return s.GetToken(SqlParserEQ, 0)
}

func (s *PropertyContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *SqlParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_property)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Identifier()
	}
	{
		p.SetState(124)
		p.Match(SqlParserEQ)
	}
	{
		p.SetState(125)
		p.Expression()
	}

	return localctx
}

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLimit returns the limit token.
	GetLimit() antlr.Token

	// SetLimit sets the limit token.
	SetLimit(antlr.Token)

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	limit  antlr.Token
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_query
	return p
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) GetLimit() antlr.Token { return s.limit }

func (s *QueryContext) SetLimit(v antlr.Token) { s.limit = v }

func (s *QueryContext) QueryTerm() IQueryTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryTermContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryTermContext)
}

func (s *QueryContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *QueryContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *QueryContext) AllSortItem() []ISortItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISortItemContext)(nil)).Elem())
	var tst = make([]ISortItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISortItemContext)
		}
	}

	return tst
}

func (s *QueryContext) SortItem(i int) ISortItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISortItemContext)
}

func (s *QueryContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *QueryContext) INTEGER_VALUE() antlr.TerminalNode {
	return s.GetToken(SqlParserINTEGER_VALUE, 0)
}

func (s *QueryContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (p *SqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.queryTerm(0)
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserORDER {
		{
			p.SetState(128)
			p.Match(SqlParserORDER)
		}
		{
			p.SetState(129)
			p.Match(SqlParserBY)
		}
		{
			p.SetState(130)
			p.SortItem()
		}
		p.SetState(135)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(131)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(132)
				p.SortItem()
			}

			p.SetState(137)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLIMIT {
		{
			p.SetState(140)
			p.Match(SqlParserLIMIT)
		}
		{
			p.SetState(141)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*QueryContext).limit = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserALL || _la == SqlParserINTEGER_VALUE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*QueryContext).limit = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IQueryTermContext is an interface to support dynamic dispatch.
type IQueryTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IQueryTermContext

	// GetRight returns the right rule contexts.
	GetRight() IQueryTermContext

	// SetLeft sets the left rule contexts.
	SetLeft(IQueryTermContext)

	// SetRight sets the right rule contexts.
	SetRight(IQueryTermContext)

	// IsQueryTermContext differentiates from other interfaces.
	IsQueryTermContext()
}

type QueryTermContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     IQueryTermContext
	operator antlr.Token
	right    IQueryTermContext
}

func NewEmptyQueryTermContext() *QueryTermContext {
	var p = new(QueryTermContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_queryTerm
	return p
}

func (*QueryTermContext) IsQueryTermContext() {}

func NewQueryTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryTermContext {
	var p = new(QueryTermContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_queryTerm

	return p
}

func (s *QueryTermContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryTermContext) GetOperator() antlr.Token { return s.operator }

func (s *QueryTermContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *QueryTermContext) GetLeft() IQueryTermContext { return s.left }

func (s *QueryTermContext) GetRight() IQueryTermContext { return s.right }

func (s *QueryTermContext) SetLeft(v IQueryTermContext) { s.left = v }

func (s *QueryTermContext) SetRight(v IQueryTermContext) { s.right = v }

func (s *QueryTermContext) QueryPrimary() IQueryPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryPrimaryContext)
}

func (s *QueryTermContext) AllQueryTerm() []IQueryTermContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IQueryTermContext)(nil)).Elem())
	var tst = make([]IQueryTermContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IQueryTermContext)
		}
	}

	return tst
}

func (s *QueryTermContext) QueryTerm(i int) IQueryTermContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryTermContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IQueryTermContext)
}

func (s *QueryTermContext) INTERSECT() antlr.TerminalNode {
	return s.GetToken(SqlParserINTERSECT, 0)
}

func (s *QueryTermContext) SetQuantifier() ISetQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *QueryTermContext) UNION() antlr.TerminalNode {
	return s.GetToken(SqlParserUNION, 0)
}

func (s *QueryTermContext) EXCEPT() antlr.TerminalNode {
	return s.GetToken(SqlParserEXCEPT, 0)
}

func (s *QueryTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterQueryTerm(s)
	}
}

func (s *QueryTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitQueryTerm(s)
	}
}

func (p *SqlParser) QueryTerm() (localctx IQueryTermContext) {
	return p.queryTerm(0)
}

func (p *SqlParser) queryTerm(_p int) (localctx IQueryTermContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewQueryTermContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IQueryTermContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, SqlParserRULE_queryTerm, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(145)
		p.QueryPrimary()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewQueryTermContext(p, _parentctx, _parentState)
				localctx.(*QueryTermContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_queryTerm)
				p.SetState(147)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(148)

					var _m = p.Match(SqlParserINTERSECT)

					localctx.(*QueryTermContext).operator = _m
				}
				p.SetState(150)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SqlParserALL || _la == SqlParserDISTINCT {
					{
						p.SetState(149)
						p.SetQuantifier()
					}

				}
				{
					p.SetState(152)

					var _x = p.queryTerm(3)

					localctx.(*QueryTermContext).right = _x
				}

			case 2:
				localctx = NewQueryTermContext(p, _parentctx, _parentState)
				localctx.(*QueryTermContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_queryTerm)
				p.SetState(153)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(154)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*QueryTermContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserEXCEPT || _la == SqlParserUNION) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*QueryTermContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(156)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == SqlParserALL || _la == SqlParserDISTINCT {
					{
						p.SetState(155)
						p.SetQuantifier()
					}

				}
				{
					p.SetState(158)

					var _x = p.queryTerm(2)

					localctx.(*QueryTermContext).right = _x
				}

			}

		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IQueryPrimaryContext is an interface to support dynamic dispatch.
type IQueryPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQueryPrimaryContext differentiates from other interfaces.
	IsQueryPrimaryContext()
}

type QueryPrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryPrimaryContext() *QueryPrimaryContext {
	var p = new(QueryPrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_queryPrimary
	return p
}

func (*QueryPrimaryContext) IsQueryPrimaryContext() {}

func NewQueryPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryPrimaryContext {
	var p = new(QueryPrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_queryPrimary

	return p
}

func (s *QueryPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryPrimaryContext) QuerySpecification() IQuerySpecificationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQuerySpecificationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQuerySpecificationContext)
}

func (s *QueryPrimaryContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *QueryPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterQueryPrimary(s)
	}
}

func (s *QueryPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitQueryPrimary(s)
	}
}

func (p *SqlParser) QueryPrimary() (localctx IQueryPrimaryContext) {
	localctx = NewQueryPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SqlParserRULE_queryPrimary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSELECT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.QuerySpecification()
		}

	case SqlParserT__0:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(166)
			p.Query()
		}
		{
			p.SetState(167)
			p.Match(SqlParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISortItemContext is an interface to support dynamic dispatch.
type ISortItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrdering returns the ordering token.
	GetOrdering() antlr.Token

	// GetNullOrdering returns the nullOrdering token.
	GetNullOrdering() antlr.Token

	// SetOrdering sets the ordering token.
	SetOrdering(antlr.Token)

	// SetNullOrdering sets the nullOrdering token.
	SetNullOrdering(antlr.Token)

	// IsSortItemContext differentiates from other interfaces.
	IsSortItemContext()
}

type SortItemContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	ordering     antlr.Token
	nullOrdering antlr.Token
}

func NewEmptySortItemContext() *SortItemContext {
	var p = new(SortItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_sortItem
	return p
}

func (*SortItemContext) IsSortItemContext() {}

func NewSortItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SortItemContext {
	var p = new(SortItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_sortItem

	return p
}

func (s *SortItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SortItemContext) GetOrdering() antlr.Token { return s.ordering }

func (s *SortItemContext) GetNullOrdering() antlr.Token { return s.nullOrdering }

func (s *SortItemContext) SetOrdering(v antlr.Token) { s.ordering = v }

func (s *SortItemContext) SetNullOrdering(v antlr.Token) { s.nullOrdering = v }

func (s *SortItemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SortItemContext) NULLS() antlr.TerminalNode {
	return s.GetToken(SqlParserNULLS, 0)
}

func (s *SortItemContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *SortItemContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *SortItemContext) FIRST() antlr.TerminalNode {
	return s.GetToken(SqlParserFIRST, 0)
}

func (s *SortItemContext) LAST() antlr.TerminalNode {
	return s.GetToken(SqlParserLAST, 0)
}

func (s *SortItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SortItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SortItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSortItem(s)
	}
}

func (s *SortItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSortItem(s)
	}
}

func (p *SqlParser) SortItem() (localctx ISortItemContext) {
	localctx = NewSortItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_sortItem)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Expression()
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserASC || _la == SqlParserDESC {
		{
			p.SetState(172)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SortItemContext).ordering = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserASC || _la == SqlParserDESC) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SortItemContext).ordering = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserNULLS {
		{
			p.SetState(175)
			p.Match(SqlParserNULLS)
		}
		{
			p.SetState(176)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*SortItemContext).nullOrdering = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserFIRST || _la == SqlParserLAST) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*SortItemContext).nullOrdering = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// IQuerySpecificationContext is an interface to support dynamic dispatch.
type IQuerySpecificationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetWhere returns the where rule contexts.
	GetWhere() IBooleanExpressionContext

	// GetHaving returns the having rule contexts.
	GetHaving() IBooleanExpressionContext

	// SetWhere sets the where rule contexts.
	SetWhere(IBooleanExpressionContext)

	// SetHaving sets the having rule contexts.
	SetHaving(IBooleanExpressionContext)

	// IsQuerySpecificationContext differentiates from other interfaces.
	IsQuerySpecificationContext()
}

type QuerySpecificationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	where  IBooleanExpressionContext
	having IBooleanExpressionContext
}

func NewEmptyQuerySpecificationContext() *QuerySpecificationContext {
	var p = new(QuerySpecificationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_querySpecification
	return p
}

func (*QuerySpecificationContext) IsQuerySpecificationContext() {}

func NewQuerySpecificationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuerySpecificationContext {
	var p = new(QuerySpecificationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_querySpecification

	return p
}

func (s *QuerySpecificationContext) GetParser() antlr.Parser { return s.parser }

func (s *QuerySpecificationContext) GetWhere() IBooleanExpressionContext { return s.where }

func (s *QuerySpecificationContext) GetHaving() IBooleanExpressionContext { return s.having }

func (s *QuerySpecificationContext) SetWhere(v IBooleanExpressionContext) { s.where = v }

func (s *QuerySpecificationContext) SetHaving(v IBooleanExpressionContext) { s.having = v }

func (s *QuerySpecificationContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *QuerySpecificationContext) AllSelectItem() []ISelectItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectItemContext)(nil)).Elem())
	var tst = make([]ISelectItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectItemContext)
		}
	}

	return tst
}

func (s *QuerySpecificationContext) SelectItem(i int) ISelectItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectItemContext)
}

func (s *QuerySpecificationContext) SetQuantifier() ISetQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *QuerySpecificationContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *QuerySpecificationContext) AllRelation() []IRelationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationContext)(nil)).Elem())
	var tst = make([]IRelationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationContext)
		}
	}

	return tst
}

func (s *QuerySpecificationContext) Relation(i int) IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *QuerySpecificationContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
}

func (s *QuerySpecificationContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SqlParserGROUP, 0)
}

func (s *QuerySpecificationContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *QuerySpecificationContext) GroupBy() IGroupByContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupByContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupByContext)
}

func (s *QuerySpecificationContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SqlParserHAVING, 0)
}

func (s *QuerySpecificationContext) AllBooleanExpression() []IBooleanExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem())
	var tst = make([]IBooleanExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBooleanExpressionContext)
		}
	}

	return tst
}

func (s *QuerySpecificationContext) BooleanExpression(i int) IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *QuerySpecificationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuerySpecificationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuerySpecificationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterQuerySpecification(s)
	}
}

func (s *QuerySpecificationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitQuerySpecification(s)
	}
}

func (p *SqlParser) QuerySpecification() (localctx IQuerySpecificationContext) {
	localctx = NewQuerySpecificationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_querySpecification)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(SqlParserSELECT)
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(180)
			p.SetQuantifier()
		}

	}
	{
		p.SetState(183)
		p.SelectItem()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(184)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(185)
				p.SelectItem()
			}

		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(191)
			p.Match(SqlParserFROM)
		}
		{
			p.SetState(192)
			p.relation(0)
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(193)
					p.Match(SqlParserT__1)
				}
				{
					p.SetState(194)
					p.relation(0)
				}

			}
			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
		}

	}
	p.SetState(204)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(202)
			p.Match(SqlParserWHERE)
		}
		{
			p.SetState(203)

			var _x = p.booleanExpression(0)

			localctx.(*QuerySpecificationContext).where = _x
		}

	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(206)
			p.Match(SqlParserGROUP)
		}
		{
			p.SetState(207)
			p.Match(SqlParserBY)
		}
		{
			p.SetState(208)
			p.GroupBy()
		}

	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(211)
			p.Match(SqlParserHAVING)
		}
		{
			p.SetState(212)

			var _x = p.booleanExpression(0)

			localctx.(*QuerySpecificationContext).having = _x
		}

	}

	return localctx
}

// IGroupByContext is an interface to support dynamic dispatch.
type IGroupByContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByContext differentiates from other interfaces.
	IsGroupByContext()
}

type GroupByContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByContext() *GroupByContext {
	var p = new(GroupByContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupBy
	return p
}

func (*GroupByContext) IsGroupByContext() {}

func NewGroupByContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByContext {
	var p = new(GroupByContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupBy

	return p
}

func (s *GroupByContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByContext) AllGroupingElement() []IGroupingElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroupingElementContext)(nil)).Elem())
	var tst = make([]IGroupingElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroupingElementContext)
		}
	}

	return tst
}

func (s *GroupByContext) GroupingElement(i int) IGroupingElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupingElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroupingElementContext)
}

func (s *GroupByContext) SetQuantifier() ISetQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *GroupByContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupBy(s)
	}
}

func (s *GroupByContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupBy(s)
	}
}

func (p *SqlParser) GroupBy() (localctx IGroupByContext) {
	localctx = NewGroupByContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SqlParserRULE_groupBy)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(215)
			p.SetQuantifier()
		}

	}
	{
		p.SetState(218)
		p.GroupingElement()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(219)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(220)
				p.GroupingElement()
			}

		}
		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IGroupingElementContext is an interface to support dynamic dispatch.
type IGroupingElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupingElementContext differentiates from other interfaces.
	IsGroupingElementContext()
}

type GroupingElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupingElementContext() *GroupingElementContext {
	var p = new(GroupingElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupingElement
	return p
}

func (*GroupingElementContext) IsGroupingElementContext() {}

func NewGroupingElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupingElementContext {
	var p = new(GroupingElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupingElement

	return p
}

func (s *GroupingElementContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupingElementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupingElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupingElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupingElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterGroupingElement(s)
	}
}

func (s *GroupingElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitGroupingElement(s)
	}
}

func (p *SqlParser) GroupingElement() (localctx IGroupingElementContext) {
	localctx = NewGroupingElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_groupingElement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Expression()
	}

	return localctx
}

// ISetQuantifierContext is an interface to support dynamic dispatch.
type ISetQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetQuantifierContext differentiates from other interfaces.
	IsSetQuantifierContext()
}

type SetQuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetQuantifierContext() *SetQuantifierContext {
	var p = new(SetQuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_setQuantifier
	return p
}

func (*SetQuantifierContext) IsSetQuantifierContext() {}

func NewSetQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetQuantifierContext {
	var p = new(SetQuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_setQuantifier

	return p
}

func (s *SetQuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SetQuantifierContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *SetQuantifierContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *SetQuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetQuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetQuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSetQuantifier(s)
	}
}

func (s *SetQuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSetQuantifier(s)
	}
}

func (p *SqlParser) SetQuantifier() (localctx ISetQuantifierContext) {
	localctx = NewSetQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_setQuantifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserALL || _la == SqlParserDISTINCT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISelectItemContext is an interface to support dynamic dispatch.
type ISelectItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectItemContext differentiates from other interfaces.
	IsSelectItemContext()
}

type SelectItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectItemContext() *SelectItemContext {
	var p = new(SelectItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectItem
	return p
}

func (*SelectItemContext) IsSelectItemContext() {}

func NewSelectItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectItemContext {
	var p = new(SelectItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectItem

	return p
}

func (s *SelectItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectItemContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SelectItemContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SelectItemContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SelectItemContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *SelectItemContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SqlParserASTERISK, 0)
}

func (s *SelectItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSelectItem(s)
	}
}

func (s *SelectItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSelectItem(s)
	}
}

func (p *SqlParser) SelectItem() (localctx ISelectItemContext) {
	localctx = NewSelectItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_selectItem)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(230)
			p.Expression()
		}
		p.SetState(235)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
			p.SetState(232)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == SqlParserAS {
				{
					p.SetState(231)
					p.Match(SqlParserAS)
				}

			}
			{
				p.SetState(234)
				p.Identifier()
			}

		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(237)
			p.QualifiedName()
		}
		{
			p.SetState(238)
			p.Match(SqlParserT__3)
		}
		{
			p.SetState(239)
			p.Match(SqlParserASTERISK)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(241)
			p.Match(SqlParserASTERISK)
		}

	}

	return localctx
}

// IRelationContext is an interface to support dynamic dispatch.
type IRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IRelationContext

	// GetRightRelation returns the rightRelation rule contexts.
	GetRightRelation() IRelationContext

	// SetLeft sets the left rule contexts.
	SetLeft(IRelationContext)

	// SetRightRelation sets the rightRelation rule contexts.
	SetRightRelation(IRelationContext)

	// IsRelationContext differentiates from other interfaces.
	IsRelationContext()
}

type RelationContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	left          IRelationContext
	rightRelation IRelationContext
}

func NewEmptyRelationContext() *RelationContext {
	var p = new(RelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_relation
	return p
}

func (*RelationContext) IsRelationContext() {}

func NewRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationContext {
	var p = new(RelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_relation

	return p
}

func (s *RelationContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationContext) GetLeft() IRelationContext { return s.left }

func (s *RelationContext) GetRightRelation() IRelationContext { return s.rightRelation }

func (s *RelationContext) SetLeft(v IRelationContext) { s.left = v }

func (s *RelationContext) SetRightRelation(v IRelationContext) { s.rightRelation = v }

func (s *RelationContext) SampledRelation() ISampledRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISampledRelationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISampledRelationContext)
}

func (s *RelationContext) JoinType() IJoinTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJoinTypeContext)
}

func (s *RelationContext) JOIN() antlr.TerminalNode {
	return s.GetToken(SqlParserJOIN, 0)
}

func (s *RelationContext) JoinCriteria() IJoinCriteriaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IJoinCriteriaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IJoinCriteriaContext)
}

func (s *RelationContext) AllRelation() []IRelationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRelationContext)(nil)).Elem())
	var tst = make([]IRelationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRelationContext)
		}
	}

	return tst
}

func (s *RelationContext) Relation(i int) IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *RelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterRelation(s)
	}
}

func (s *RelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitRelation(s)
	}
}

func (p *SqlParser) Relation() (localctx IRelationContext) {
	return p.relation(0)
}

func (p *SqlParser) relation(_p int) (localctx IRelationContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRelationContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRelationContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, SqlParserRULE_relation, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		p.SampledRelation()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRelationContext(p, _parentctx, _parentState)
			localctx.(*RelationContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_relation)
			p.SetState(247)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(248)
				p.JoinType()
			}
			{
				p.SetState(249)
				p.Match(SqlParserJOIN)
			}
			{
				p.SetState(250)

				var _x = p.relation(0)

				localctx.(*RelationContext).rightRelation = _x
			}
			{
				p.SetState(251)
				p.JoinCriteria()
			}

		}
		p.SetState(257)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext())
	}

	return localctx
}

// IJoinTypeContext is an interface to support dynamic dispatch.
type IJoinTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinTypeContext differentiates from other interfaces.
	IsJoinTypeContext()
}

type JoinTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinTypeContext() *JoinTypeContext {
	var p = new(JoinTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_joinType
	return p
}

func (*JoinTypeContext) IsJoinTypeContext() {}

func NewJoinTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinTypeContext {
	var p = new(JoinTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_joinType

	return p
}

func (s *JoinTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinTypeContext) INNER() antlr.TerminalNode {
	return s.GetToken(SqlParserINNER, 0)
}

func (s *JoinTypeContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SqlParserLEFT, 0)
}

func (s *JoinTypeContext) OUTER() antlr.TerminalNode {
	return s.GetToken(SqlParserOUTER, 0)
}

func (s *JoinTypeContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(SqlParserRIGHT, 0)
}

func (s *JoinTypeContext) FULL() antlr.TerminalNode {
	return s.GetToken(SqlParserFULL, 0)
}

func (s *JoinTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterJoinType(s)
	}
}

func (s *JoinTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitJoinType(s)
	}
}

func (p *SqlParser) JoinType() (localctx IJoinTypeContext) {
	localctx = NewJoinTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_joinType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(273)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserINNER, SqlParserJOIN:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserINNER {
			{
				p.SetState(258)
				p.Match(SqlParserINNER)
			}

		}

	case SqlParserLEFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)
			p.Match(SqlParserLEFT)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserOUTER {
			{
				p.SetState(262)
				p.Match(SqlParserOUTER)
			}

		}

	case SqlParserRIGHT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(265)
			p.Match(SqlParserRIGHT)
		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserOUTER {
			{
				p.SetState(266)
				p.Match(SqlParserOUTER)
			}

		}

	case SqlParserFULL:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(269)
			p.Match(SqlParserFULL)
		}
		p.SetState(271)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserOUTER {
			{
				p.SetState(270)
				p.Match(SqlParserOUTER)
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IJoinCriteriaContext is an interface to support dynamic dispatch.
type IJoinCriteriaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinCriteriaContext differentiates from other interfaces.
	IsJoinCriteriaContext()
}

type JoinCriteriaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinCriteriaContext() *JoinCriteriaContext {
	var p = new(JoinCriteriaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_joinCriteria
	return p
}

func (*JoinCriteriaContext) IsJoinCriteriaContext() {}

func NewJoinCriteriaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinCriteriaContext {
	var p = new(JoinCriteriaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_joinCriteria

	return p
}

func (s *JoinCriteriaContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinCriteriaContext) ON() antlr.TerminalNode {
	return s.GetToken(SqlParserON, 0)
}

func (s *JoinCriteriaContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *JoinCriteriaContext) USING() antlr.TerminalNode {
	return s.GetToken(SqlParserUSING, 0)
}

func (s *JoinCriteriaContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *JoinCriteriaContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *JoinCriteriaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinCriteriaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinCriteriaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterJoinCriteria(s)
	}
}

func (s *JoinCriteriaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitJoinCriteria(s)
	}
}

func (p *SqlParser) JoinCriteria() (localctx IJoinCriteriaContext) {
	localctx = NewJoinCriteriaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_joinCriteria)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(289)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(SqlParserON)
		}
		{
			p.SetState(276)
			p.booleanExpression(0)
		}

	case SqlParserUSING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(277)
			p.Match(SqlParserUSING)
		}
		{
			p.SetState(278)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(279)
			p.Identifier()
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(280)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(281)
				p.Identifier()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(287)
			p.Match(SqlParserT__2)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISampleTypeContext is an interface to support dynamic dispatch.
type ISampleTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSampleTypeContext differentiates from other interfaces.
	IsSampleTypeContext()
}

type SampleTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySampleTypeContext() *SampleTypeContext {
	var p = new(SampleTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_sampleType
	return p
}

func (*SampleTypeContext) IsSampleTypeContext() {}

func NewSampleTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SampleTypeContext {
	var p = new(SampleTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_sampleType

	return p
}

func (s *SampleTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *SampleTypeContext) BERNOULLI() antlr.TerminalNode {
	return s.GetToken(SqlParserBERNOULLI, 0)
}

func (s *SampleTypeContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(SqlParserSYSTEM, 0)
}

func (s *SampleTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SampleTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SampleTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSampleType(s)
	}
}

func (s *SampleTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSampleType(s)
	}
}

func (p *SqlParser) SampleType() (localctx ISampleTypeContext) {
	localctx = NewSampleTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_sampleType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserBERNOULLI || _la == SqlParserSYSTEM) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISampledRelationContext is an interface to support dynamic dispatch.
type ISampledRelationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSampledRelationContext differentiates from other interfaces.
	IsSampledRelationContext()
}

type SampledRelationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySampledRelationContext() *SampledRelationContext {
	var p = new(SampledRelationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_sampledRelation
	return p
}

func (*SampledRelationContext) IsSampledRelationContext() {}

func NewSampledRelationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SampledRelationContext {
	var p = new(SampledRelationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_sampledRelation

	return p
}

func (s *SampledRelationContext) GetParser() antlr.Parser { return s.parser }

func (s *SampledRelationContext) RelationPrimary() IRelationPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationPrimaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationPrimaryContext)
}

func (s *SampledRelationContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *SampledRelationContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SampledRelationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SampledRelationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SampledRelationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterSampledRelation(s)
	}
}

func (s *SampledRelationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitSampledRelation(s)
	}
}

func (p *SqlParser) SampledRelation() (localctx ISampledRelationContext) {
	localctx = NewSampledRelationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_sampledRelation)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.RelationPrimary()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		p.SetState(295)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserAS {
			{
				p.SetState(294)
				p.Match(SqlParserAS)
			}

		}
		{
			p.SetState(297)
			p.Identifier()
		}

	}

	return localctx
}

// IRelationPrimaryContext is an interface to support dynamic dispatch.
type IRelationPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRelationPrimaryContext differentiates from other interfaces.
	IsRelationPrimaryContext()
}

type RelationPrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationPrimaryContext() *RelationPrimaryContext {
	var p = new(RelationPrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_relationPrimary
	return p
}

func (*RelationPrimaryContext) IsRelationPrimaryContext() {}

func NewRelationPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationPrimaryContext {
	var p = new(RelationPrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_relationPrimary

	return p
}

func (s *RelationPrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationPrimaryContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *RelationPrimaryContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *RelationPrimaryContext) Relation() IRelationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRelationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRelationContext)
}

func (s *RelationPrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationPrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationPrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterRelationPrimary(s)
	}
}

func (s *RelationPrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitRelationPrimary(s)
	}
}

func (p *SqlParser) RelationPrimary() (localctx IRelationPrimaryContext) {
	localctx = NewRelationPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SqlParserRULE_relationPrimary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(300)
			p.QualifiedName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(301)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(302)
			p.Query()
		}
		{
			p.SetState(303)
			p.Match(SqlParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(305)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(306)
			p.relation(0)
		}
		{
			p.SetState(307)
			p.Match(SqlParserT__2)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *SqlParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SqlParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.booleanExpression(0)
	}

	return localctx
}

// IBooleanExpressionContext is an interface to support dynamic dispatch.
type IBooleanExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IBooleanExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IBooleanExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IBooleanExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IBooleanExpressionContext)

	// IsBooleanExpressionContext differentiates from other interfaces.
	IsBooleanExpressionContext()
}

type BooleanExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     IBooleanExpressionContext
	operator antlr.Token
	right    IBooleanExpressionContext
}

func NewEmptyBooleanExpressionContext() *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_booleanExpression
	return p
}

func (*BooleanExpressionContext) IsBooleanExpressionContext() {}

func NewBooleanExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanExpressionContext {
	var p = new(BooleanExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_booleanExpression

	return p
}

func (s *BooleanExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanExpressionContext) GetOperator() antlr.Token { return s.operator }

func (s *BooleanExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *BooleanExpressionContext) GetLeft() IBooleanExpressionContext { return s.left }

func (s *BooleanExpressionContext) GetRight() IBooleanExpressionContext { return s.right }

func (s *BooleanExpressionContext) SetLeft(v IBooleanExpressionContext) { s.left = v }

func (s *BooleanExpressionContext) SetRight(v IBooleanExpressionContext) { s.right = v }

func (s *BooleanExpressionContext) Predicated() IPredicatedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicatedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicatedContext)
}

func (s *BooleanExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *BooleanExpressionContext) AllBooleanExpression() []IBooleanExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem())
	var tst = make([]IBooleanExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBooleanExpressionContext)
		}
	}

	return tst
}

func (s *BooleanExpressionContext) BooleanExpression(i int) IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *BooleanExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *BooleanExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *BooleanExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterBooleanExpression(s)
	}
}

func (s *BooleanExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitBooleanExpression(s)
	}
}

func (p *SqlParser) BooleanExpression() (localctx IBooleanExpressionContext) {
	return p.booleanExpression(0)
}

func (p *SqlParser) booleanExpression(_p int) (localctx IBooleanExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBooleanExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBooleanExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, SqlParserRULE_booleanExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserT__0, SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFALSE, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULL, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRUE, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserPLUS, SqlParserMINUS, SqlParserSTRING, SqlParserINTEGER_VALUE, SqlParserDOUBLE_VALUE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER:
		{
			p.SetState(314)
			p.Predicated()
		}

	case SqlParserNOT:
		{
			p.SetState(315)
			p.Match(SqlParserNOT)
		}
		{
			p.SetState(316)
			p.booleanExpression(3)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(325)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBooleanExpressionContext(p, _parentctx, _parentState)
				localctx.(*BooleanExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_booleanExpression)
				p.SetState(319)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(320)

					var _m = p.Match(SqlParserAND)

					localctx.(*BooleanExpressionContext).operator = _m
				}
				{
					p.SetState(321)

					var _x = p.booleanExpression(3)

					localctx.(*BooleanExpressionContext).right = _x
				}

			case 2:
				localctx = NewBooleanExpressionContext(p, _parentctx, _parentState)
				localctx.(*BooleanExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_booleanExpression)
				p.SetState(322)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(323)

					var _m = p.Match(SqlParserOR)

					localctx.(*BooleanExpressionContext).operator = _m
				}
				{
					p.SetState(324)

					var _x = p.booleanExpression(2)

					localctx.(*BooleanExpressionContext).right = _x
				}

			}

		}
		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicatedContext is an interface to support dynamic dispatch.
type IPredicatedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicatedContext differentiates from other interfaces.
	IsPredicatedContext()
}

type PredicatedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicatedContext() *PredicatedContext {
	var p = new(PredicatedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_predicated
	return p
}

func (*PredicatedContext) IsPredicatedContext() {}

func NewPredicatedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicatedContext {
	var p = new(PredicatedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_predicated

	return p
}

func (s *PredicatedContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicatedContext) ValueExpression() IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *PredicatedContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicatedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicatedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicatedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterPredicated(s)
	}
}

func (s *PredicatedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitPredicated(s)
	}
}

func (p *SqlParser) Predicated() (localctx IPredicatedContext) {
	localctx = NewPredicatedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_predicated)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(330)
		p.valueExpression(0)
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(331)
			p.Predicate()
		}

	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetRight returns the right rule contexts.
	GetRight() IValueExpressionContext

	// GetLower returns the lower rule contexts.
	GetLower() IValueExpressionContext

	// GetUpper returns the upper rule contexts.
	GetUpper() IValueExpressionContext

	// GetPattern returns the pattern rule contexts.
	GetPattern() IValueExpressionContext

	// GetEscape returns the escape rule contexts.
	GetEscape() IValueExpressionContext

	// SetRight sets the right rule contexts.
	SetRight(IValueExpressionContext)

	// SetLower sets the lower rule contexts.
	SetLower(IValueExpressionContext)

	// SetUpper sets the upper rule contexts.
	SetUpper(IValueExpressionContext)

	// SetPattern sets the pattern rule contexts.
	SetPattern(IValueExpressionContext)

	// SetEscape sets the escape rule contexts.
	SetEscape(IValueExpressionContext)

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	right   IValueExpressionContext
	lower   IValueExpressionContext
	upper   IValueExpressionContext
	pattern IValueExpressionContext
	escape  IValueExpressionContext
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) GetRight() IValueExpressionContext { return s.right }

func (s *PredicateContext) GetLower() IValueExpressionContext { return s.lower }

func (s *PredicateContext) GetUpper() IValueExpressionContext { return s.upper }

func (s *PredicateContext) GetPattern() IValueExpressionContext { return s.pattern }

func (s *PredicateContext) GetEscape() IValueExpressionContext { return s.escape }

func (s *PredicateContext) SetRight(v IValueExpressionContext) { s.right = v }

func (s *PredicateContext) SetLower(v IValueExpressionContext) { s.lower = v }

func (s *PredicateContext) SetUpper(v IValueExpressionContext) { s.upper = v }

func (s *PredicateContext) SetPattern(v IValueExpressionContext) { s.pattern = v }

func (s *PredicateContext) SetEscape(v IValueExpressionContext) { s.escape = v }

func (s *PredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *PredicateContext) AllValueExpression() []IValueExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem())
	var tst = make([]IValueExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueExpressionContext)
		}
	}

	return tst
}

func (s *PredicateContext) ValueExpression(i int) IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *PredicateContext) ComparisonQuantifier() IComparisonQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonQuantifierContext)
}

func (s *PredicateContext) Query() IQueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQueryContext)
}

func (s *PredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SqlParserBETWEEN, 0)
}

func (s *PredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *PredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *PredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(SqlParserIN, 0)
}

func (s *PredicateContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PredicateContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserLIKE, 0)
}

func (s *PredicateContext) ESCAPE() antlr.TerminalNode {
	return s.GetToken(SqlParserESCAPE, 0)
}

func (s *PredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(SqlParserIS, 0)
}

func (s *PredicateContext) NULL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL, 0)
}

func (s *PredicateContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *PredicateContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *SqlParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SqlParserRULE_predicate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)
			p.ComparisonOperator()
		}
		{
			p.SetState(335)

			var _x = p.valueExpression(0)

			localctx.(*PredicateContext).right = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.ComparisonOperator()
		}
		{
			p.SetState(338)
			p.ComparisonQuantifier()
		}
		{
			p.SetState(339)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(340)
			p.Query()
		}
		{
			p.SetState(341)
			p.Match(SqlParserT__2)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(343)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(346)
			p.Match(SqlParserBETWEEN)
		}
		{
			p.SetState(347)

			var _x = p.valueExpression(0)

			localctx.(*PredicateContext).lower = _x
		}
		{
			p.SetState(348)
			p.Match(SqlParserAND)
		}
		{
			p.SetState(349)

			var _x = p.valueExpression(0)

			localctx.(*PredicateContext).upper = _x
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(351)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(354)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(355)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(356)
			p.Expression()
		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(357)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(358)
				p.Expression()
			}

			p.SetState(363)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(364)
			p.Match(SqlParserT__2)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(366)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(369)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(370)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(371)
			p.Query()
		}
		{
			p.SetState(372)
			p.Match(SqlParserT__2)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(374)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(377)
			p.Match(SqlParserLIKE)
		}
		{
			p.SetState(378)

			var _x = p.valueExpression(0)

			localctx.(*PredicateContext).pattern = _x
		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(379)
				p.Match(SqlParserESCAPE)
			}
			{
				p.SetState(380)

				var _x = p.valueExpression(0)

				localctx.(*PredicateContext).escape = _x
			}

		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(383)
			p.Match(SqlParserIS)
		}
		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(384)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(387)
			p.Match(SqlParserNULL)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(388)
			p.Match(SqlParserIS)
		}
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(389)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(392)
			p.Match(SqlParserDISTINCT)
		}
		{
			p.SetState(393)
			p.Match(SqlParserFROM)
		}
		{
			p.SetState(394)

			var _x = p.valueExpression(0)

			localctx.(*PredicateContext).right = _x
		}

	}

	return localctx
}

// IValueExpressionContext is an interface to support dynamic dispatch.
type IValueExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOperator returns the operator token.
	GetOperator() antlr.Token

	// SetOperator sets the operator token.
	SetOperator(antlr.Token)

	// GetLeft returns the left rule contexts.
	GetLeft() IValueExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IValueExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IValueExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IValueExpressionContext)

	// IsValueExpressionContext differentiates from other interfaces.
	IsValueExpressionContext()
}

type ValueExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	left     IValueExpressionContext
	operator antlr.Token
	right    IValueExpressionContext
}

func NewEmptyValueExpressionContext() *ValueExpressionContext {
	var p = new(ValueExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_valueExpression
	return p
}

func (*ValueExpressionContext) IsValueExpressionContext() {}

func NewValueExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_valueExpression

	return p
}

func (s *ValueExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExpressionContext) GetOperator() antlr.Token { return s.operator }

func (s *ValueExpressionContext) SetOperator(v antlr.Token) { s.operator = v }

func (s *ValueExpressionContext) GetLeft() IValueExpressionContext { return s.left }

func (s *ValueExpressionContext) GetRight() IValueExpressionContext { return s.right }

func (s *ValueExpressionContext) SetLeft(v IValueExpressionContext) { s.left = v }

func (s *ValueExpressionContext) SetRight(v IValueExpressionContext) { s.right = v }

func (s *ValueExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ValueExpressionContext) AllValueExpression() []IValueExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem())
	var tst = make([]IValueExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueExpressionContext)
		}
	}

	return tst
}

func (s *ValueExpressionContext) ValueExpression(i int) IValueExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueExpressionContext)
}

func (s *ValueExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SqlParserMINUS, 0)
}

func (s *ValueExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SqlParserPLUS, 0)
}

func (s *ValueExpressionContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SqlParserASTERISK, 0)
}

func (s *ValueExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SqlParserSLASH, 0)
}

func (s *ValueExpressionContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(SqlParserPERCENT, 0)
}

func (s *ValueExpressionContext) CONCAT() antlr.TerminalNode {
	return s.GetToken(SqlParserCONCAT, 0)
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterValueExpression(s)
	}
}

func (s *ValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitValueExpression(s)
	}
}

func (p *SqlParser) ValueExpression() (localctx IValueExpressionContext) {
	return p.valueExpression(0)
}

func (p *SqlParser) valueExpression(_p int) (localctx IValueExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewValueExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IValueExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 54
	p.EnterRecursionRule(localctx, 54, SqlParserRULE_valueExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserT__0, SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFALSE, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULL, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRUE, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserSTRING, SqlParserINTEGER_VALUE, SqlParserDOUBLE_VALUE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER:
		{
			p.SetState(398)
			p.primaryExpression(0)
		}

	case SqlParserPLUS, SqlParserMINUS:
		{
			p.SetState(399)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*ValueExpressionContext).operator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserPLUS || _la == SqlParserMINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*ValueExpressionContext).operator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(400)
			p.valueExpression(4)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(412)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
			case 1:
				localctx = NewValueExpressionContext(p, _parentctx, _parentState)
				localctx.(*ValueExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_valueExpression)
				p.SetState(403)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(404)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ValueExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-193)&-(0x1f+1)) == 0 && ((1<<uint((_la-193)))&((1<<(SqlParserASTERISK-193))|(1<<(SqlParserSLASH-193))|(1<<(SqlParserPERCENT-193)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ValueExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(405)

					var _x = p.valueExpression(4)

					localctx.(*ValueExpressionContext).right = _x
				}

			case 2:
				localctx = NewValueExpressionContext(p, _parentctx, _parentState)
				localctx.(*ValueExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_valueExpression)
				p.SetState(406)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(407)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ValueExpressionContext).operator = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == SqlParserPLUS || _la == SqlParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ValueExpressionContext).operator = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(408)

					var _x = p.valueExpression(3)

					localctx.(*ValueExpressionContext).right = _x
				}

			case 3:
				localctx = NewValueExpressionContext(p, _parentctx, _parentState)
				localctx.(*ValueExpressionContext).left = _prevctx
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_valueExpression)
				p.SetState(409)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(410)
					p.Match(SqlParserCONCAT)
				}
				{
					p.SetState(411)

					var _x = p.valueExpression(2)

					localctx.(*ValueExpressionContext).right = _x
				}

			}

		}
		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetBase returns the base rule contexts.
	GetBase() IPrimaryExpressionContext

	// GetFieldName returns the fieldName rule contexts.
	GetFieldName() IIdentifierContext

	// SetBase sets the base rule contexts.
	SetBase(IPrimaryExpressionContext)

	// SetFieldName sets the fieldName rule contexts.
	SetFieldName(IIdentifierContext)

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	base      IPrimaryExpressionContext
	fieldName IIdentifierContext
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_primaryExpression
	return p
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) GetBase() IPrimaryExpressionContext { return s.base }

func (s *PrimaryExpressionContext) GetFieldName() IIdentifierContext { return s.fieldName }

func (s *PrimaryExpressionContext) SetBase(v IPrimaryExpressionContext) { s.base = v }

func (s *PrimaryExpressionContext) SetFieldName(v IIdentifierContext) { s.fieldName = v }

func (s *PrimaryExpressionContext) NULL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL, 0)
}

func (s *PrimaryExpressionContext) Number() INumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumberContext)
}

func (s *PrimaryExpressionContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *PrimaryExpressionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *PrimaryExpressionContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PrimaryExpressionContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *PrimaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *PrimaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *PrimaryExpressionContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *PrimaryExpressionContext) AllSortItem() []ISortItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISortItemContext)(nil)).Elem())
	var tst = make([]ISortItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISortItemContext)
		}
	}

	return tst
}

func (s *PrimaryExpressionContext) SortItem(i int) ISortItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISortItemContext)
}

func (s *PrimaryExpressionContext) SetQuantifier() ISetQuantifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISetQuantifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISetQuantifierContext)
}

func (s *PrimaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *SqlParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	return p.primaryExpression(0)
}

func (p *SqlParser) primaryExpression(_p int) (localctx IPrimaryExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IPrimaryExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, SqlParserRULE_primaryExpression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(418)
			p.Match(SqlParserNULL)
		}

	case 2:
		{
			p.SetState(419)
			p.Number()
		}

	case 3:
		{
			p.SetState(420)
			p.BooleanValue()
		}

	case 4:
		{
			p.SetState(421)
			p.StringValue()
		}

	case 5:
		{
			p.SetState(422)
			p.Identifier()
		}

	case 6:
		{
			p.SetState(423)
			p.QualifiedName()
		}
		{
			p.SetState(424)
			p.Match(SqlParserT__0)
		}
		p.SetState(436)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SqlParserT__0)|(1<<SqlParserADD)|(1<<SqlParserALL)|(1<<SqlParserANALYZE)|(1<<SqlParserANY)|(1<<SqlParserARRAY)|(1<<SqlParserASC)|(1<<SqlParserAT)|(1<<SqlParserBERNOULLI)|(1<<SqlParserCALL)|(1<<SqlParserCASCADE)|(1<<SqlParserCATALOGS)|(1<<SqlParserCOALESCE)|(1<<SqlParserCOLUMN)|(1<<SqlParserCOLUMNS)|(1<<SqlParserCOMMENT)|(1<<SqlParserCOMMIT)|(1<<SqlParserCOMMITTED))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SqlParserCURRENT-33))|(1<<(SqlParserDATA-33))|(1<<(SqlParserDATE-33))|(1<<(SqlParserDAY-33))|(1<<(SqlParserDESC-33))|(1<<(SqlParserDISTINCT-33))|(1<<(SqlParserDISTRIBUTED-33))|(1<<(SqlParserEXCLUDING-33))|(1<<(SqlParserEXPLAIN-33))|(1<<(SqlParserFALSE-33))|(1<<(SqlParserFILTER-33))|(1<<(SqlParserFIRST-33))|(1<<(SqlParserFOLLOWING-33))|(1<<(SqlParserFORMAT-33))|(1<<(SqlParserFUNCTIONS-33)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(SqlParserGRANT-65))|(1<<(SqlParserGRANTS-65))|(1<<(SqlParserGRAPHVIZ-65))|(1<<(SqlParserHOUR-65))|(1<<(SqlParserIF-65))|(1<<(SqlParserINCLUDING-65))|(1<<(SqlParserINPUT-65))|(1<<(SqlParserINTEGER-65))|(1<<(SqlParserINTERVAL-65))|(1<<(SqlParserISOLATION-65))|(1<<(SqlParserLAST-65))|(1<<(SqlParserLATERAL-65))|(1<<(SqlParserLEVEL-65))|(1<<(SqlParserLIMIT-65))|(1<<(SqlParserLOGICAL-65))|(1<<(SqlParserMAP-65))|(1<<(SqlParserMINUTE-65))|(1<<(SqlParserMONTH-65)))) != 0) || (((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(SqlParserNFC-98))|(1<<(SqlParserNFD-98))|(1<<(SqlParserNFKC-98))|(1<<(SqlParserNFKD-98))|(1<<(SqlParserNO-98))|(1<<(SqlParserNOT-98))|(1<<(SqlParserNULL-98))|(1<<(SqlParserNULLIF-98))|(1<<(SqlParserNULLS-98))|(1<<(SqlParserONLY-98))|(1<<(SqlParserOPTION-98))|(1<<(SqlParserORDINALITY-98))|(1<<(SqlParserOUTPUT-98))|(1<<(SqlParserOVER-98))|(1<<(SqlParserPARTITION-98))|(1<<(SqlParserPARTITIONS-98))|(1<<(SqlParserPOSITION-98))|(1<<(SqlParserPRECEDING-98))|(1<<(SqlParserPRIVILEGES-98))|(1<<(SqlParserPROPERTIES-98))|(1<<(SqlParserPUBLIC-98))|(1<<(SqlParserRANGE-98))|(1<<(SqlParserREAD-98))|(1<<(SqlParserRENAME-98))|(1<<(SqlParserREPEATABLE-98)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(SqlParserREPLACE-130))|(1<<(SqlParserRESET-130))|(1<<(SqlParserRESTRICT-130))|(1<<(SqlParserREVOKE-130))|(1<<(SqlParserROLLBACK-130))|(1<<(SqlParserROW-130))|(1<<(SqlParserROWS-130))|(1<<(SqlParserSCHEMA-130))|(1<<(SqlParserSCHEMAS-130))|(1<<(SqlParserSECOND-130))|(1<<(SqlParserSERIALIZABLE-130))|(1<<(SqlParserSESSION-130))|(1<<(SqlParserSET-130))|(1<<(SqlParserSETS-130))|(1<<(SqlParserSHOW-130))|(1<<(SqlParserSMALLINT-130))|(1<<(SqlParserSOME-130))|(1<<(SqlParserSTART-130))|(1<<(SqlParserSTATS-130))|(1<<(SqlParserSUBSTRING-130))|(1<<(SqlParserSYSTEM-130))|(1<<(SqlParserTABLES-130))|(1<<(SqlParserTABLESAMPLE-130))|(1<<(SqlParserTEXT-130))|(1<<(SqlParserTIME-130))|(1<<(SqlParserTIMESTAMP-130))|(1<<(SqlParserTINYINT-130)))) != 0) || (((_la-162)&-(0x1f+1)) == 0 && ((1<<uint((_la-162)))&((1<<(SqlParserTO-162))|(1<<(SqlParserTRANSACTION-162))|(1<<(SqlParserTRUE-162))|(1<<(SqlParserTRY_CAST-162))|(1<<(SqlParserTYPE-162))|(1<<(SqlParserUNBOUNDED-162))|(1<<(SqlParserUNCOMMITTED-162))|(1<<(SqlParserUSE-162))|(1<<(SqlParserVALIDATE-162))|(1<<(SqlParserVERBOSE-162))|(1<<(SqlParserVIEW-162))|(1<<(SqlParserWORK-162))|(1<<(SqlParserWRITE-162))|(1<<(SqlParserYEAR-162))|(1<<(SqlParserZONE-162))|(1<<(SqlParserPLUS-162))|(1<<(SqlParserMINUS-162)))) != 0) || (((_la-197)&-(0x1f+1)) == 0 && ((1<<uint((_la-197)))&((1<<(SqlParserSTRING-197))|(1<<(SqlParserINTEGER_VALUE-197))|(1<<(SqlParserDOUBLE_VALUE-197))|(1<<(SqlParserIDENTIFIER-197))|(1<<(SqlParserDIGIT_IDENTIFIER-197))|(1<<(SqlParserQUOTED_IDENTIFIER-197)))) != 0) {
			p.SetState(426)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(425)
					p.SetQuantifier()
				}

			}
			{
				p.SetState(428)
				p.Expression()
			}
			p.SetState(433)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SqlParserT__1 {
				{
					p.SetState(429)
					p.Match(SqlParserT__1)
				}
				{
					p.SetState(430)
					p.Expression()
				}

				p.SetState(435)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		p.SetState(448)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserORDER {
			{
				p.SetState(438)
				p.Match(SqlParserORDER)
			}
			{
				p.SetState(439)
				p.Match(SqlParserBY)
			}
			{
				p.SetState(440)
				p.SortItem()
			}
			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SqlParserT__1 {
				{
					p.SetState(441)
					p.Match(SqlParserT__1)
				}
				{
					p.SetState(442)
					p.SortItem()
				}

				p.SetState(447)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}

		}
		{
			p.SetState(450)
			p.Match(SqlParserT__2)
		}

	case 7:
		{
			p.SetState(452)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(453)
			p.Expression()
		}
		{
			p.SetState(454)
			p.Match(SqlParserT__2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(463)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewPrimaryExpressionContext(p, _parentctx, _parentState)
			localctx.(*PrimaryExpressionContext).base = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_primaryExpression)
			p.SetState(458)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(459)
				p.Match(SqlParserT__3)
			}
			{
				p.SetState(460)

				var _x = p.Identifier()

				localctx.(*PrimaryExpressionContext).fieldName = _x
			}

		}
		p.SetState(465)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (p *SqlParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SqlParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(466)
		p.Match(SqlParserSTRING)
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SqlParserEQ, 0)
}

func (s *ComparisonOperatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SqlParserNEQ, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SqlParserLT, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(SqlParserLTE, 0)
}

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SqlParserGT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(SqlParserGTE, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (p *SqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SqlParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-185)&-(0x1f+1)) == 0 && ((1<<uint((_la-185)))&((1<<(SqlParserEQ-185))|(1<<(SqlParserNEQ-185))|(1<<(SqlParserLT-185))|(1<<(SqlParserLTE-185))|(1<<(SqlParserGT-185))|(1<<(SqlParserGTE-185)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonQuantifierContext is an interface to support dynamic dispatch.
type IComparisonQuantifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonQuantifierContext differentiates from other interfaces.
	IsComparisonQuantifierContext()
}

type ComparisonQuantifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonQuantifierContext() *ComparisonQuantifierContext {
	var p = new(ComparisonQuantifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_comparisonQuantifier
	return p
}

func (*ComparisonQuantifierContext) IsComparisonQuantifierContext() {}

func NewComparisonQuantifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonQuantifierContext {
	var p = new(ComparisonQuantifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonQuantifier

	return p
}

func (s *ComparisonQuantifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonQuantifierContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *ComparisonQuantifierContext) SOME() antlr.TerminalNode {
	return s.GetToken(SqlParserSOME, 0)
}

func (s *ComparisonQuantifierContext) ANY() antlr.TerminalNode {
	return s.GetToken(SqlParserANY, 0)
}

func (s *ComparisonQuantifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonQuantifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonQuantifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterComparisonQuantifier(s)
	}
}

func (s *ComparisonQuantifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitComparisonQuantifier(s)
	}
}

func (p *SqlParser) ComparisonQuantifier() (localctx IComparisonQuantifierContext) {
	localctx = NewComparisonQuantifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SqlParserRULE_comparisonQuantifier)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserALL || _la == SqlParserANY || _la == SqlParserSOME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SqlParserTRUE, 0)
}

func (s *BooleanValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SqlParserFALSE, 0)
}

func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (p *SqlParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SqlParserRULE_booleanValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(472)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserFALSE || _la == SqlParserTRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITypeSqlContext is an interface to support dynamic dispatch.
type ITypeSqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSqlContext differentiates from other interfaces.
	IsTypeSqlContext()
}

type TypeSqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSqlContext() *TypeSqlContext {
	var p = new(TypeSqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_typeSql
	return p
}

func (*TypeSqlContext) IsTypeSqlContext() {}

func NewTypeSqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSqlContext {
	var p = new(TypeSqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_typeSql

	return p
}

func (s *TypeSqlContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSqlContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(SqlParserARRAY, 0)
}

func (s *TypeSqlContext) AllTypeSql() []ITypeSqlContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSqlContext)(nil)).Elem())
	var tst = make([]ITypeSqlContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSqlContext)
		}
	}

	return tst
}

func (s *TypeSqlContext) TypeSql(i int) ITypeSqlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSqlContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSqlContext)
}

func (s *TypeSqlContext) MAP() antlr.TerminalNode {
	return s.GetToken(SqlParserMAP, 0)
}

func (s *TypeSqlContext) ROW() antlr.TerminalNode {
	return s.GetToken(SqlParserROW, 0)
}

func (s *TypeSqlContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *TypeSqlContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TypeSqlContext) BaseType() IBaseTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseTypeContext)
}

func (s *TypeSqlContext) AllTypeParameter() []ITypeParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeParameterContext)(nil)).Elem())
	var tst = make([]ITypeParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParameterContext)
		}
	}

	return tst
}

func (s *TypeSqlContext) TypeParameter(i int) ITypeParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParameterContext)
}

func (s *TypeSqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTypeSql(s)
	}
}

func (s *TypeSqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTypeSql(s)
	}
}

func (p *SqlParser) TypeSql() (localctx ITypeSqlContext) {
	return p.typeSql(0)
}

func (p *SqlParser) typeSql(_p int) (localctx ITypeSqlContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeSqlContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeSqlContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, SqlParserRULE_typeSql, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(475)
			p.Match(SqlParserARRAY)
		}
		{
			p.SetState(476)
			p.Match(SqlParserLT)
		}
		{
			p.SetState(477)
			p.typeSql(0)
		}
		{
			p.SetState(478)
			p.Match(SqlParserGT)
		}

	case 2:
		{
			p.SetState(480)
			p.Match(SqlParserMAP)
		}
		{
			p.SetState(481)
			p.Match(SqlParserLT)
		}
		{
			p.SetState(482)
			p.typeSql(0)
		}
		{
			p.SetState(483)
			p.Match(SqlParserT__1)
		}
		{
			p.SetState(484)
			p.typeSql(0)
		}
		{
			p.SetState(485)
			p.Match(SqlParserGT)
		}

	case 3:
		{
			p.SetState(487)
			p.Match(SqlParserROW)
		}
		{
			p.SetState(488)
			p.Match(SqlParserT__0)
		}
		{
			p.SetState(489)
			p.Identifier()
		}
		{
			p.SetState(490)
			p.typeSql(0)
		}
		p.SetState(497)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(491)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(492)
				p.Identifier()
			}
			{
				p.SetState(493)
				p.typeSql(0)
			}

			p.SetState(499)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(500)
			p.Match(SqlParserT__2)
		}

	case 4:
		{
			p.SetState(502)
			p.BaseType()
		}
		p.SetState(514)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(503)
				p.Match(SqlParserT__0)
			}
			{
				p.SetState(504)
				p.TypeParameter()
			}
			p.SetState(509)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for _la == SqlParserT__1 {
				{
					p.SetState(505)
					p.Match(SqlParserT__1)
				}
				{
					p.SetState(506)
					p.TypeParameter()
				}

				p.SetState(511)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(512)
				p.Match(SqlParserT__2)
			}

		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewTypeSqlContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_typeSql)
			p.SetState(518)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(519)
				p.Match(SqlParserARRAY)
			}

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext())
	}

	return localctx
}

// ITypeParameterContext is an interface to support dynamic dispatch.
type ITypeParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeParameterContext differentiates from other interfaces.
	IsTypeParameterContext()
}

type TypeParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeParameterContext() *TypeParameterContext {
	var p = new(TypeParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_typeParameter
	return p
}

func (*TypeParameterContext) IsTypeParameterContext() {}

func NewTypeParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_typeParameter

	return p
}

func (s *TypeParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParameterContext) INTEGER_VALUE() antlr.TerminalNode {
	return s.GetToken(SqlParserINTEGER_VALUE, 0)
}

func (s *TypeParameterContext) TypeSql() ITypeSqlContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSqlContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSqlContext)
}

func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (p *SqlParser) TypeParameter() (localctx ITypeParameterContext) {
	localctx = NewTypeParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SqlParserRULE_typeParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(527)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserINTEGER_VALUE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(525)
			p.Match(SqlParserINTEGER_VALUE)
		}

	case SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER, SqlParserTIME_WITH_TIME_ZONE, SqlParserTIMESTAMP_WITH_TIME_ZONE, SqlParserDOUBLE_PRECISION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.typeSql(0)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBaseTypeContext is an interface to support dynamic dispatch.
type IBaseTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseTypeContext differentiates from other interfaces.
	IsBaseTypeContext()
}

type BaseTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseTypeContext() *BaseTypeContext {
	var p = new(BaseTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_baseType
	return p
}

func (*BaseTypeContext) IsBaseTypeContext() {}

func NewBaseTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseTypeContext {
	var p = new(BaseTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_baseType

	return p
}

func (s *BaseTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseTypeContext) TIME_WITH_TIME_ZONE() antlr.TerminalNode {
	return s.GetToken(SqlParserTIME_WITH_TIME_ZONE, 0)
}

func (s *BaseTypeContext) TIMESTAMP_WITH_TIME_ZONE() antlr.TerminalNode {
	return s.GetToken(SqlParserTIMESTAMP_WITH_TIME_ZONE, 0)
}

func (s *BaseTypeContext) DOUBLE_PRECISION() antlr.TerminalNode {
	return s.GetToken(SqlParserDOUBLE_PRECISION, 0)
}

func (s *BaseTypeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (p *SqlParser) BaseType() (localctx IBaseTypeContext) {
	localctx = NewBaseTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SqlParserRULE_baseType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(533)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserTIME_WITH_TIME_ZONE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(529)
			p.Match(SqlParserTIME_WITH_TIME_ZONE)
		}

	case SqlParserTIMESTAMP_WITH_TIME_ZONE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(530)
			p.Match(SqlParserTIMESTAMP_WITH_TIME_ZONE)
		}

	case SqlParserDOUBLE_PRECISION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(531)
			p.Match(SqlParserDOUBLE_PRECISION)
		}

	case SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(532)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IWhenClauseContext is an interface to support dynamic dispatch.
type IWhenClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCondition returns the condition rule contexts.
	GetCondition() IExpressionContext

	// GetResult returns the result rule contexts.
	GetResult() IExpressionContext

	// SetCondition sets the condition rule contexts.
	SetCondition(IExpressionContext)

	// SetResult sets the result rule contexts.
	SetResult(IExpressionContext)

	// IsWhenClauseContext differentiates from other interfaces.
	IsWhenClauseContext()
}

type WhenClauseContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	condition IExpressionContext
	result    IExpressionContext
}

func NewEmptyWhenClauseContext() *WhenClauseContext {
	var p = new(WhenClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_whenClause
	return p
}

func (*WhenClauseContext) IsWhenClauseContext() {}

func NewWhenClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhenClauseContext {
	var p = new(WhenClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_whenClause

	return p
}

func (s *WhenClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhenClauseContext) GetCondition() IExpressionContext { return s.condition }

func (s *WhenClauseContext) GetResult() IExpressionContext { return s.result }

func (s *WhenClauseContext) SetCondition(v IExpressionContext) { s.condition = v }

func (s *WhenClauseContext) SetResult(v IExpressionContext) { s.result = v }

func (s *WhenClauseContext) WHEN() antlr.TerminalNode {
	return s.GetToken(SqlParserWHEN, 0)
}

func (s *WhenClauseContext) THEN() antlr.TerminalNode {
	return s.GetToken(SqlParserTHEN, 0)
}

func (s *WhenClauseContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *WhenClauseContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhenClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhenClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhenClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterWhenClause(s)
	}
}

func (s *WhenClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitWhenClause(s)
	}
}

func (p *SqlParser) WhenClause() (localctx IWhenClauseContext) {
	localctx = NewWhenClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SqlParserRULE_whenClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(535)
		p.Match(SqlParserWHEN)
	}
	{
		p.SetState(536)

		var _x = p.Expression()

		localctx.(*WhenClauseContext).condition = _x
	}
	{
		p.SetState(537)
		p.Match(SqlParserTHEN)
	}
	{
		p.SetState(538)

		var _x = p.Expression()

		localctx.(*WhenClauseContext).result = _x
	}

	return localctx
}

// IFilterContext is an interface to support dynamic dispatch.
type IFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilterContext differentiates from other interfaces.
	IsFilterContext()
}

type FilterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilterContext() *FilterContext {
	var p = new(FilterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_filter
	return p
}

func (*FilterContext) IsFilterContext() {}

func NewFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FilterContext {
	var p = new(FilterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_filter

	return p
}

func (s *FilterContext) GetParser() antlr.Parser { return s.parser }

func (s *FilterContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SqlParserFILTER, 0)
}

func (s *FilterContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
}

func (s *FilterContext) BooleanExpression() IBooleanExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanExpressionContext)
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *SqlParser) Filter() (localctx IFilterContext) {
	localctx = NewFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SqlParserRULE_filter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.Match(SqlParserFILTER)
	}
	{
		p.SetState(541)
		p.Match(SqlParserT__0)
	}
	{
		p.SetState(542)
		p.Match(SqlParserWHERE)
	}
	{
		p.SetState(543)
		p.booleanExpression(0)
	}
	{
		p.SetState(544)
		p.Match(SqlParserT__2)
	}

	return localctx
}

// IOverContext is an interface to support dynamic dispatch.
type IOverContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPartition returns the partition rule context list.
	GetPartition() []IExpressionContext

	// SetPartition sets the partition rule context list.
	SetPartition([]IExpressionContext)

	// IsOverContext differentiates from other interfaces.
	IsOverContext()
}

type OverContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_expression IExpressionContext
	partition   []IExpressionContext
}

func NewEmptyOverContext() *OverContext {
	var p = new(OverContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_over
	return p
}

func (*OverContext) IsOverContext() {}

func NewOverContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverContext {
	var p = new(OverContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_over

	return p
}

func (s *OverContext) GetParser() antlr.Parser { return s.parser }

func (s *OverContext) Get_expression() IExpressionContext { return s._expression }

func (s *OverContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *OverContext) GetPartition() []IExpressionContext { return s.partition }

func (s *OverContext) SetPartition(v []IExpressionContext) { s.partition = v }

func (s *OverContext) OVER() antlr.TerminalNode {
	return s.GetToken(SqlParserOVER, 0)
}

func (s *OverContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(SqlParserPARTITION, 0)
}

func (s *OverContext) AllBY() []antlr.TerminalNode {
	return s.GetTokens(SqlParserBY)
}

func (s *OverContext) BY(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserBY, i)
}

func (s *OverContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *OverContext) AllSortItem() []ISortItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISortItemContext)(nil)).Elem())
	var tst = make([]ISortItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISortItemContext)
		}
	}

	return tst
}

func (s *OverContext) SortItem(i int) ISortItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISortItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISortItemContext)
}

func (s *OverContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *OverContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OverContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterOver(s)
	}
}

func (s *OverContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitOver(s)
	}
}

func (p *SqlParser) Over() (localctx IOverContext) {
	localctx = NewOverContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SqlParserRULE_over)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(SqlParserOVER)
	}
	{
		p.SetState(547)
		p.Match(SqlParserT__0)
	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserPARTITION {
		{
			p.SetState(548)
			p.Match(SqlParserPARTITION)
		}
		{
			p.SetState(549)
			p.Match(SqlParserBY)
		}
		{
			p.SetState(550)

			var _x = p.Expression()

			localctx.(*OverContext)._expression = _x
		}
		localctx.(*OverContext).partition = append(localctx.(*OverContext).partition, localctx.(*OverContext)._expression)
		p.SetState(555)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(551)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(552)

				var _x = p.Expression()

				localctx.(*OverContext)._expression = _x
			}
			localctx.(*OverContext).partition = append(localctx.(*OverContext).partition, localctx.(*OverContext)._expression)

			p.SetState(557)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	p.SetState(570)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserORDER {
		{
			p.SetState(560)
			p.Match(SqlParserORDER)
		}
		{
			p.SetState(561)
			p.Match(SqlParserBY)
		}
		{
			p.SetState(562)
			p.SortItem()
		}
		p.SetState(567)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserT__1 {
			{
				p.SetState(563)
				p.Match(SqlParserT__1)
			}
			{
				p.SetState(564)
				p.SortItem()
			}

			p.SetState(569)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(572)
		p.Match(SqlParserT__2)
	}

	return localctx
}

// IPrivilegeContext is an interface to support dynamic dispatch.
type IPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrivilegeContext differentiates from other interfaces.
	IsPrivilegeContext()
}

type PrivilegeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivilegeContext() *PrivilegeContext {
	var p = new(PrivilegeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_privilege
	return p
}

func (*PrivilegeContext) IsPrivilegeContext() {}

func NewPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivilegeContext {
	var p = new(PrivilegeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_privilege

	return p
}

func (s *PrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivilegeContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *PrivilegeContext) DELETE() antlr.TerminalNode {
	return s.GetToken(SqlParserDELETE, 0)
}

func (s *PrivilegeContext) INSERT() antlr.TerminalNode {
	return s.GetToken(SqlParserINSERT, 0)
}

func (s *PrivilegeContext) Identifier() IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *PrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterPrivilege(s)
	}
}

func (s *PrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitPrivilege(s)
	}
}

func (p *SqlParser) Privilege() (localctx IPrivilegeContext) {
	localctx = NewPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SqlParserRULE_privilege)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(578)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSELECT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(574)
			p.Match(SqlParserSELECT)
		}

	case SqlParserDELETE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(575)
			p.Match(SqlParserDELETE)
		}

	case SqlParserINSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(576)
			p.Match(SqlParserINSERT)
		}

	case SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE, SqlParserIDENTIFIER, SqlParserDIGIT_IDENTIFIER, SqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(577)
			p.Identifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) AllIdentifier() []IIdentifierContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IIdentifierContext)(nil)).Elem())
	var tst = make([]IIdentifierContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IIdentifierContext)
		}
	}

	return tst
}

func (s *QualifiedNameContext) Identifier(i int) IIdentifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentifierContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (p *SqlParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SqlParserRULE_qualifiedName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(580)
		p.Identifier()
	}
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(581)
				p.Match(SqlParserT__3)
			}
			{
				p.SetState(582)
				p.Identifier()
			}

		}
		p.SetState(587)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 72, p.GetParserRuleContext())
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SqlParserIDENTIFIER, 0)
}

func (s *IdentifierContext) QUOTED_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SqlParserQUOTED_IDENTIFIER, 0)
}

func (s *IdentifierContext) NonReserved() INonReservedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonReservedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonReservedContext)
}

func (s *IdentifierContext) DIGIT_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SqlParserDIGIT_IDENTIFIER, 0)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitIdentifier(s)
	}
}

func (p *SqlParser) Identifier() (localctx IIdentifierContext) {
	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SqlParserRULE_identifier)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(592)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(588)
			p.Match(SqlParserIDENTIFIER)
		}

	case SqlParserQUOTED_IDENTIFIER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(589)
			p.Match(SqlParserQUOTED_IDENTIFIER)
		}

	case SqlParserADD, SqlParserALL, SqlParserANALYZE, SqlParserANY, SqlParserARRAY, SqlParserASC, SqlParserAT, SqlParserBERNOULLI, SqlParserCALL, SqlParserCASCADE, SqlParserCATALOGS, SqlParserCOALESCE, SqlParserCOLUMN, SqlParserCOLUMNS, SqlParserCOMMENT, SqlParserCOMMIT, SqlParserCOMMITTED, SqlParserCURRENT, SqlParserDATA, SqlParserDATE, SqlParserDAY, SqlParserDESC, SqlParserDISTRIBUTED, SqlParserEXCLUDING, SqlParserEXPLAIN, SqlParserFILTER, SqlParserFIRST, SqlParserFOLLOWING, SqlParserFORMAT, SqlParserFUNCTIONS, SqlParserGRANT, SqlParserGRANTS, SqlParserGRAPHVIZ, SqlParserHOUR, SqlParserIF, SqlParserINCLUDING, SqlParserINPUT, SqlParserINTEGER, SqlParserINTERVAL, SqlParserISOLATION, SqlParserLAST, SqlParserLATERAL, SqlParserLEVEL, SqlParserLIMIT, SqlParserLOGICAL, SqlParserMAP, SqlParserMINUTE, SqlParserMONTH, SqlParserNFC, SqlParserNFD, SqlParserNFKC, SqlParserNFKD, SqlParserNO, SqlParserNULLIF, SqlParserNULLS, SqlParserONLY, SqlParserOPTION, SqlParserORDINALITY, SqlParserOUTPUT, SqlParserOVER, SqlParserPARTITION, SqlParserPARTITIONS, SqlParserPOSITION, SqlParserPRECEDING, SqlParserPRIVILEGES, SqlParserPROPERTIES, SqlParserPUBLIC, SqlParserRANGE, SqlParserREAD, SqlParserRENAME, SqlParserREPEATABLE, SqlParserREPLACE, SqlParserRESET, SqlParserRESTRICT, SqlParserREVOKE, SqlParserROLLBACK, SqlParserROW, SqlParserROWS, SqlParserSCHEMA, SqlParserSCHEMAS, SqlParserSECOND, SqlParserSERIALIZABLE, SqlParserSESSION, SqlParserSET, SqlParserSETS, SqlParserSHOW, SqlParserSMALLINT, SqlParserSOME, SqlParserSTART, SqlParserSTATS, SqlParserSUBSTRING, SqlParserSYSTEM, SqlParserTABLES, SqlParserTABLESAMPLE, SqlParserTEXT, SqlParserTIME, SqlParserTIMESTAMP, SqlParserTINYINT, SqlParserTO, SqlParserTRANSACTION, SqlParserTRY_CAST, SqlParserTYPE, SqlParserUNBOUNDED, SqlParserUNCOMMITTED, SqlParserUSE, SqlParserVALIDATE, SqlParserVERBOSE, SqlParserVIEW, SqlParserWORK, SqlParserWRITE, SqlParserYEAR, SqlParserZONE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(590)
			p.NonReserved()
		}

	case SqlParserDIGIT_IDENTIFIER:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(591)
			p.Match(SqlParserDIGIT_IDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberContext is an interface to support dynamic dispatch.
type INumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberContext differentiates from other interfaces.
	IsNumberContext()
}

type NumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberContext() *NumberContext {
	var p = new(NumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_number
	return p
}

func (*NumberContext) IsNumberContext() {}

func NewNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberContext {
	var p = new(NumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_number

	return p
}

func (s *NumberContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberContext) DOUBLE_VALUE() antlr.TerminalNode {
	return s.GetToken(SqlParserDOUBLE_VALUE, 0)
}

func (s *NumberContext) INTEGER_VALUE() antlr.TerminalNode {
	return s.GetToken(SqlParserINTEGER_VALUE, 0)
}

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitNumber(s)
	}
}

func (p *SqlParser) Number() (localctx INumberContext) {
	localctx = NewNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SqlParserRULE_number)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(594)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserINTEGER_VALUE || _la == SqlParserDOUBLE_VALUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INonReservedContext is an interface to support dynamic dispatch.
type INonReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonReservedContext differentiates from other interfaces.
	IsNonReservedContext()
}

type NonReservedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonReservedContext() *NonReservedContext {
	var p = new(NonReservedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_nonReserved
	return p
}

func (*NonReservedContext) IsNonReservedContext() {}

func NewNonReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonReservedContext {
	var p = new(NonReservedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_nonReserved

	return p
}

func (s *NonReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *NonReservedContext) ADD() antlr.TerminalNode {
	return s.GetToken(SqlParserADD, 0)
}

func (s *NonReservedContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *NonReservedContext) ANALYZE() antlr.TerminalNode {
	return s.GetToken(SqlParserANALYZE, 0)
}

func (s *NonReservedContext) ANY() antlr.TerminalNode {
	return s.GetToken(SqlParserANY, 0)
}

func (s *NonReservedContext) ARRAY() antlr.TerminalNode {
	return s.GetToken(SqlParserARRAY, 0)
}

func (s *NonReservedContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *NonReservedContext) AT() antlr.TerminalNode {
	return s.GetToken(SqlParserAT, 0)
}

func (s *NonReservedContext) BERNOULLI() antlr.TerminalNode {
	return s.GetToken(SqlParserBERNOULLI, 0)
}

func (s *NonReservedContext) CALL() antlr.TerminalNode {
	return s.GetToken(SqlParserCALL, 0)
}

func (s *NonReservedContext) CASCADE() antlr.TerminalNode {
	return s.GetToken(SqlParserCASCADE, 0)
}

func (s *NonReservedContext) CATALOGS() antlr.TerminalNode {
	return s.GetToken(SqlParserCATALOGS, 0)
}

func (s *NonReservedContext) COALESCE() antlr.TerminalNode {
	return s.GetToken(SqlParserCOALESCE, 0)
}

func (s *NonReservedContext) COLUMN() antlr.TerminalNode {
	return s.GetToken(SqlParserCOLUMN, 0)
}

func (s *NonReservedContext) COLUMNS() antlr.TerminalNode {
	return s.GetToken(SqlParserCOLUMNS, 0)
}

func (s *NonReservedContext) COMMENT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMENT, 0)
}

func (s *NonReservedContext) COMMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMIT, 0)
}

func (s *NonReservedContext) COMMITTED() antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMITTED, 0)
}

func (s *NonReservedContext) CURRENT() antlr.TerminalNode {
	return s.GetToken(SqlParserCURRENT, 0)
}

func (s *NonReservedContext) DATA() antlr.TerminalNode {
	return s.GetToken(SqlParserDATA, 0)
}

func (s *NonReservedContext) DATE() antlr.TerminalNode {
	return s.GetToken(SqlParserDATE, 0)
}

func (s *NonReservedContext) DAY() antlr.TerminalNode {
	return s.GetToken(SqlParserDAY, 0)
}

func (s *NonReservedContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *NonReservedContext) DISTRIBUTED() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTRIBUTED, 0)
}

func (s *NonReservedContext) EXCLUDING() antlr.TerminalNode {
	return s.GetToken(SqlParserEXCLUDING, 0)
}

func (s *NonReservedContext) EXPLAIN() antlr.TerminalNode {
	return s.GetToken(SqlParserEXPLAIN, 0)
}

func (s *NonReservedContext) FILTER() antlr.TerminalNode {
	return s.GetToken(SqlParserFILTER, 0)
}

func (s *NonReservedContext) FIRST() antlr.TerminalNode {
	return s.GetToken(SqlParserFIRST, 0)
}

func (s *NonReservedContext) FOLLOWING() antlr.TerminalNode {
	return s.GetToken(SqlParserFOLLOWING, 0)
}

func (s *NonReservedContext) FORMAT() antlr.TerminalNode {
	return s.GetToken(SqlParserFORMAT, 0)
}

func (s *NonReservedContext) FUNCTIONS() antlr.TerminalNode {
	return s.GetToken(SqlParserFUNCTIONS, 0)
}

func (s *NonReservedContext) GRANT() antlr.TerminalNode {
	return s.GetToken(SqlParserGRANT, 0)
}

func (s *NonReservedContext) GRANTS() antlr.TerminalNode {
	return s.GetToken(SqlParserGRANTS, 0)
}

func (s *NonReservedContext) GRAPHVIZ() antlr.TerminalNode {
	return s.GetToken(SqlParserGRAPHVIZ, 0)
}

func (s *NonReservedContext) HOUR() antlr.TerminalNode {
	return s.GetToken(SqlParserHOUR, 0)
}

func (s *NonReservedContext) IF() antlr.TerminalNode {
	return s.GetToken(SqlParserIF, 0)
}

func (s *NonReservedContext) INCLUDING() antlr.TerminalNode {
	return s.GetToken(SqlParserINCLUDING, 0)
}

func (s *NonReservedContext) INPUT() antlr.TerminalNode {
	return s.GetToken(SqlParserINPUT, 0)
}

func (s *NonReservedContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SqlParserINTEGER, 0)
}

func (s *NonReservedContext) INTERVAL() antlr.TerminalNode {
	return s.GetToken(SqlParserINTERVAL, 0)
}

func (s *NonReservedContext) ISOLATION() antlr.TerminalNode {
	return s.GetToken(SqlParserISOLATION, 0)
}

func (s *NonReservedContext) LAST() antlr.TerminalNode {
	return s.GetToken(SqlParserLAST, 0)
}

func (s *NonReservedContext) LATERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserLATERAL, 0)
}

func (s *NonReservedContext) LEVEL() antlr.TerminalNode {
	return s.GetToken(SqlParserLEVEL, 0)
}

func (s *NonReservedContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *NonReservedContext) LOGICAL() antlr.TerminalNode {
	return s.GetToken(SqlParserLOGICAL, 0)
}

func (s *NonReservedContext) MAP() antlr.TerminalNode {
	return s.GetToken(SqlParserMAP, 0)
}

func (s *NonReservedContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(SqlParserMINUTE, 0)
}

func (s *NonReservedContext) MONTH() antlr.TerminalNode {
	return s.GetToken(SqlParserMONTH, 0)
}

func (s *NonReservedContext) NFC() antlr.TerminalNode {
	return s.GetToken(SqlParserNFC, 0)
}

func (s *NonReservedContext) NFD() antlr.TerminalNode {
	return s.GetToken(SqlParserNFD, 0)
}

func (s *NonReservedContext) NFKC() antlr.TerminalNode {
	return s.GetToken(SqlParserNFKC, 0)
}

func (s *NonReservedContext) NFKD() antlr.TerminalNode {
	return s.GetToken(SqlParserNFKD, 0)
}

func (s *NonReservedContext) NO() antlr.TerminalNode {
	return s.GetToken(SqlParserNO, 0)
}

func (s *NonReservedContext) NULLIF() antlr.TerminalNode {
	return s.GetToken(SqlParserNULLIF, 0)
}

func (s *NonReservedContext) NULLS() antlr.TerminalNode {
	return s.GetToken(SqlParserNULLS, 0)
}

func (s *NonReservedContext) ONLY() antlr.TerminalNode {
	return s.GetToken(SqlParserONLY, 0)
}

func (s *NonReservedContext) OPTION() antlr.TerminalNode {
	return s.GetToken(SqlParserOPTION, 0)
}

func (s *NonReservedContext) ORDINALITY() antlr.TerminalNode {
	return s.GetToken(SqlParserORDINALITY, 0)
}

func (s *NonReservedContext) OUTPUT() antlr.TerminalNode {
	return s.GetToken(SqlParserOUTPUT, 0)
}

func (s *NonReservedContext) OVER() antlr.TerminalNode {
	return s.GetToken(SqlParserOVER, 0)
}

func (s *NonReservedContext) PARTITION() antlr.TerminalNode {
	return s.GetToken(SqlParserPARTITION, 0)
}

func (s *NonReservedContext) PARTITIONS() antlr.TerminalNode {
	return s.GetToken(SqlParserPARTITIONS, 0)
}

func (s *NonReservedContext) POSITION() antlr.TerminalNode {
	return s.GetToken(SqlParserPOSITION, 0)
}

func (s *NonReservedContext) PRECEDING() antlr.TerminalNode {
	return s.GetToken(SqlParserPRECEDING, 0)
}

func (s *NonReservedContext) PRIVILEGES() antlr.TerminalNode {
	return s.GetToken(SqlParserPRIVILEGES, 0)
}

func (s *NonReservedContext) PROPERTIES() antlr.TerminalNode {
	return s.GetToken(SqlParserPROPERTIES, 0)
}

func (s *NonReservedContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(SqlParserPUBLIC, 0)
}

func (s *NonReservedContext) RANGE() antlr.TerminalNode {
	return s.GetToken(SqlParserRANGE, 0)
}

func (s *NonReservedContext) READ() antlr.TerminalNode {
	return s.GetToken(SqlParserREAD, 0)
}

func (s *NonReservedContext) RENAME() antlr.TerminalNode {
	return s.GetToken(SqlParserRENAME, 0)
}

func (s *NonReservedContext) REPEATABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserREPEATABLE, 0)
}

func (s *NonReservedContext) REPLACE() antlr.TerminalNode {
	return s.GetToken(SqlParserREPLACE, 0)
}

func (s *NonReservedContext) RESET() antlr.TerminalNode {
	return s.GetToken(SqlParserRESET, 0)
}

func (s *NonReservedContext) RESTRICT() antlr.TerminalNode {
	return s.GetToken(SqlParserRESTRICT, 0)
}

func (s *NonReservedContext) REVOKE() antlr.TerminalNode {
	return s.GetToken(SqlParserREVOKE, 0)
}

func (s *NonReservedContext) ROLLBACK() antlr.TerminalNode {
	return s.GetToken(SqlParserROLLBACK, 0)
}

func (s *NonReservedContext) ROW() antlr.TerminalNode {
	return s.GetToken(SqlParserROW, 0)
}

func (s *NonReservedContext) ROWS() antlr.TerminalNode {
	return s.GetToken(SqlParserROWS, 0)
}

func (s *NonReservedContext) SCHEMA() antlr.TerminalNode {
	return s.GetToken(SqlParserSCHEMA, 0)
}

func (s *NonReservedContext) SCHEMAS() antlr.TerminalNode {
	return s.GetToken(SqlParserSCHEMAS, 0)
}

func (s *NonReservedContext) SECOND() antlr.TerminalNode {
	return s.GetToken(SqlParserSECOND, 0)
}

func (s *NonReservedContext) SERIALIZABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserSERIALIZABLE, 0)
}

func (s *NonReservedContext) SESSION() antlr.TerminalNode {
	return s.GetToken(SqlParserSESSION, 0)
}

func (s *NonReservedContext) SET() antlr.TerminalNode {
	return s.GetToken(SqlParserSET, 0)
}

func (s *NonReservedContext) SETS() antlr.TerminalNode {
	return s.GetToken(SqlParserSETS, 0)
}

func (s *NonReservedContext) SHOW() antlr.TerminalNode {
	return s.GetToken(SqlParserSHOW, 0)
}

func (s *NonReservedContext) SMALLINT() antlr.TerminalNode {
	return s.GetToken(SqlParserSMALLINT, 0)
}

func (s *NonReservedContext) SOME() antlr.TerminalNode {
	return s.GetToken(SqlParserSOME, 0)
}

func (s *NonReservedContext) START() antlr.TerminalNode {
	return s.GetToken(SqlParserSTART, 0)
}

func (s *NonReservedContext) STATS() antlr.TerminalNode {
	return s.GetToken(SqlParserSTATS, 0)
}

func (s *NonReservedContext) SUBSTRING() antlr.TerminalNode {
	return s.GetToken(SqlParserSUBSTRING, 0)
}

func (s *NonReservedContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(SqlParserSYSTEM, 0)
}

func (s *NonReservedContext) TABLES() antlr.TerminalNode {
	return s.GetToken(SqlParserTABLES, 0)
}

func (s *NonReservedContext) TABLESAMPLE() antlr.TerminalNode {
	return s.GetToken(SqlParserTABLESAMPLE, 0)
}

func (s *NonReservedContext) TEXT() antlr.TerminalNode {
	return s.GetToken(SqlParserTEXT, 0)
}

func (s *NonReservedContext) TIME() antlr.TerminalNode {
	return s.GetToken(SqlParserTIME, 0)
}

func (s *NonReservedContext) TIMESTAMP() antlr.TerminalNode {
	return s.GetToken(SqlParserTIMESTAMP, 0)
}

func (s *NonReservedContext) TINYINT() antlr.TerminalNode {
	return s.GetToken(SqlParserTINYINT, 0)
}

func (s *NonReservedContext) TO() antlr.TerminalNode {
	return s.GetToken(SqlParserTO, 0)
}

func (s *NonReservedContext) TRANSACTION() antlr.TerminalNode {
	return s.GetToken(SqlParserTRANSACTION, 0)
}

func (s *NonReservedContext) TRY_CAST() antlr.TerminalNode {
	return s.GetToken(SqlParserTRY_CAST, 0)
}

func (s *NonReservedContext) TYPE() antlr.TerminalNode {
	return s.GetToken(SqlParserTYPE, 0)
}

func (s *NonReservedContext) UNBOUNDED() antlr.TerminalNode {
	return s.GetToken(SqlParserUNBOUNDED, 0)
}

func (s *NonReservedContext) UNCOMMITTED() antlr.TerminalNode {
	return s.GetToken(SqlParserUNCOMMITTED, 0)
}

func (s *NonReservedContext) USE() antlr.TerminalNode {
	return s.GetToken(SqlParserUSE, 0)
}

func (s *NonReservedContext) VALIDATE() antlr.TerminalNode {
	return s.GetToken(SqlParserVALIDATE, 0)
}

func (s *NonReservedContext) VERBOSE() antlr.TerminalNode {
	return s.GetToken(SqlParserVERBOSE, 0)
}

func (s *NonReservedContext) VIEW() antlr.TerminalNode {
	return s.GetToken(SqlParserVIEW, 0)
}

func (s *NonReservedContext) WORK() antlr.TerminalNode {
	return s.GetToken(SqlParserWORK, 0)
}

func (s *NonReservedContext) WRITE() antlr.TerminalNode {
	return s.GetToken(SqlParserWRITE, 0)
}

func (s *NonReservedContext) YEAR() antlr.TerminalNode {
	return s.GetToken(SqlParserYEAR, 0)
}

func (s *NonReservedContext) ZONE() antlr.TerminalNode {
	return s.GetToken(SqlParserZONE, 0)
}

func (s *NonReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonReservedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.EnterNonReserved(s)
	}
}

func (s *NonReservedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlListener); ok {
		listenerT.ExitNonReserved(s)
	}
}

func (p *SqlParser) NonReserved() (localctx INonReservedContext) {
	localctx = NewNonReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SqlParserRULE_nonReserved)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SqlParserADD)|(1<<SqlParserALL)|(1<<SqlParserANALYZE)|(1<<SqlParserANY)|(1<<SqlParserARRAY)|(1<<SqlParserASC)|(1<<SqlParserAT)|(1<<SqlParserBERNOULLI)|(1<<SqlParserCALL)|(1<<SqlParserCASCADE)|(1<<SqlParserCATALOGS)|(1<<SqlParserCOALESCE)|(1<<SqlParserCOLUMN)|(1<<SqlParserCOLUMNS)|(1<<SqlParserCOMMENT)|(1<<SqlParserCOMMIT)|(1<<SqlParserCOMMITTED))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SqlParserCURRENT-33))|(1<<(SqlParserDATA-33))|(1<<(SqlParserDATE-33))|(1<<(SqlParserDAY-33))|(1<<(SqlParserDESC-33))|(1<<(SqlParserDISTRIBUTED-33))|(1<<(SqlParserEXCLUDING-33))|(1<<(SqlParserEXPLAIN-33))|(1<<(SqlParserFILTER-33))|(1<<(SqlParserFIRST-33))|(1<<(SqlParserFOLLOWING-33))|(1<<(SqlParserFORMAT-33))|(1<<(SqlParserFUNCTIONS-33)))) != 0) || (((_la-65)&-(0x1f+1)) == 0 && ((1<<uint((_la-65)))&((1<<(SqlParserGRANT-65))|(1<<(SqlParserGRANTS-65))|(1<<(SqlParserGRAPHVIZ-65))|(1<<(SqlParserHOUR-65))|(1<<(SqlParserIF-65))|(1<<(SqlParserINCLUDING-65))|(1<<(SqlParserINPUT-65))|(1<<(SqlParserINTEGER-65))|(1<<(SqlParserINTERVAL-65))|(1<<(SqlParserISOLATION-65))|(1<<(SqlParserLAST-65))|(1<<(SqlParserLATERAL-65))|(1<<(SqlParserLEVEL-65))|(1<<(SqlParserLIMIT-65))|(1<<(SqlParserLOGICAL-65))|(1<<(SqlParserMAP-65))|(1<<(SqlParserMINUTE-65))|(1<<(SqlParserMONTH-65)))) != 0) || (((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(SqlParserNFC-98))|(1<<(SqlParserNFD-98))|(1<<(SqlParserNFKC-98))|(1<<(SqlParserNFKD-98))|(1<<(SqlParserNO-98))|(1<<(SqlParserNULLIF-98))|(1<<(SqlParserNULLS-98))|(1<<(SqlParserONLY-98))|(1<<(SqlParserOPTION-98))|(1<<(SqlParserORDINALITY-98))|(1<<(SqlParserOUTPUT-98))|(1<<(SqlParserOVER-98))|(1<<(SqlParserPARTITION-98))|(1<<(SqlParserPARTITIONS-98))|(1<<(SqlParserPOSITION-98))|(1<<(SqlParserPRECEDING-98))|(1<<(SqlParserPRIVILEGES-98))|(1<<(SqlParserPROPERTIES-98))|(1<<(SqlParserPUBLIC-98))|(1<<(SqlParserRANGE-98))|(1<<(SqlParserREAD-98))|(1<<(SqlParserRENAME-98))|(1<<(SqlParserREPEATABLE-98)))) != 0) || (((_la-130)&-(0x1f+1)) == 0 && ((1<<uint((_la-130)))&((1<<(SqlParserREPLACE-130))|(1<<(SqlParserRESET-130))|(1<<(SqlParserRESTRICT-130))|(1<<(SqlParserREVOKE-130))|(1<<(SqlParserROLLBACK-130))|(1<<(SqlParserROW-130))|(1<<(SqlParserROWS-130))|(1<<(SqlParserSCHEMA-130))|(1<<(SqlParserSCHEMAS-130))|(1<<(SqlParserSECOND-130))|(1<<(SqlParserSERIALIZABLE-130))|(1<<(SqlParserSESSION-130))|(1<<(SqlParserSET-130))|(1<<(SqlParserSETS-130))|(1<<(SqlParserSHOW-130))|(1<<(SqlParserSMALLINT-130))|(1<<(SqlParserSOME-130))|(1<<(SqlParserSTART-130))|(1<<(SqlParserSTATS-130))|(1<<(SqlParserSUBSTRING-130))|(1<<(SqlParserSYSTEM-130))|(1<<(SqlParserTABLES-130))|(1<<(SqlParserTABLESAMPLE-130))|(1<<(SqlParserTEXT-130))|(1<<(SqlParserTIME-130))|(1<<(SqlParserTIMESTAMP-130))|(1<<(SqlParserTINYINT-130)))) != 0) || (((_la-162)&-(0x1f+1)) == 0 && ((1<<uint((_la-162)))&((1<<(SqlParserTO-162))|(1<<(SqlParserTRANSACTION-162))|(1<<(SqlParserTRY_CAST-162))|(1<<(SqlParserTYPE-162))|(1<<(SqlParserUNBOUNDED-162))|(1<<(SqlParserUNCOMMITTED-162))|(1<<(SqlParserUSE-162))|(1<<(SqlParserVALIDATE-162))|(1<<(SqlParserVERBOSE-162))|(1<<(SqlParserVIEW-162))|(1<<(SqlParserWORK-162))|(1<<(SqlParserWRITE-162))|(1<<(SqlParserYEAR-162))|(1<<(SqlParserZONE-162)))) != 0)) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *QueryTermContext = nil
		if localctx != nil {
			t = localctx.(*QueryTermContext)
		}
		return p.QueryTerm_Sempred(t, predIndex)

	case 17:
		var t *RelationContext = nil
		if localctx != nil {
			t = localctx.(*RelationContext)
		}
		return p.Relation_Sempred(t, predIndex)

	case 24:
		var t *BooleanExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BooleanExpressionContext)
		}
		return p.BooleanExpression_Sempred(t, predIndex)

	case 27:
		var t *ValueExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ValueExpressionContext)
		}
		return p.ValueExpression_Sempred(t, predIndex)

	case 28:
		var t *PrimaryExpressionContext = nil
		if localctx != nil {
			t = localctx.(*PrimaryExpressionContext)
		}
		return p.PrimaryExpression_Sempred(t, predIndex)

	case 33:
		var t *TypeSqlContext = nil
		if localctx != nil {
			t = localctx.(*TypeSqlContext)
		}
		return p.TypeSql_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SqlParser) QueryTerm_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) Relation_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) BooleanExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) ValueExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) PrimaryExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 8:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) TypeSql_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 9:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
