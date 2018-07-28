package HiveConnector

var MD_SQL = `
select name, type from
(
select C.COLUMN_NAME as name, C.TYPE_NAME as type, C.INTEGER_IDX as sortkey from
TBLS as T
join COLUMNS_V2 as C on C.CD_ID=T.TBL_ID
join DBS as D on D.DB_ID=T.DB_ID
where D.NAME='%s' and T.TBL_NAME='%s'

union all

select P.PKEY_NAME as name, P.PKEY_TYPE as type, P.INTEGER_IDX+1000000 as sortkey from
TBLS as T
join DBS as D on D.DB_ID=T.DB_ID
join PARTITION_KEYS as P on T.TBL_ID=P.TBL_ID
where D.NAME='%s' and  T.TBL_NAME='%s'

order by sortkey
) t;
`

var PARTITION_MD_SQL = `
select P.PKEY_NAME as name, P.PKEY_TYPE as type from 
TBLS as T 
join DBS as D on D.DB_ID=T.DB_ID
join PARTITION_KEYS as P on T.TBL_ID=P.TBL_ID
where D.NAME='%s' and T.TBL_NAME='%s'
order by P.INTEGER_IDX;
`

var PARTITION_DATA_SQL = `
select SDS.LOCATION as location, SDS.INPUT_FORMAT as filetype, PKV.PART_KEY_VAL as pkey from
TBLS as T 
join PARTITIONS as P on T.TBL_ID=P.TBL_ID
join PARTITION_KEY_VALS as PKV on P.PART_ID=PKV.PART_ID
join SDS as SDS on SDS.SD_ID=P.SD_ID
join DBS as D on D.DB_ID=T.DB_ID
where D.NAME='%s' and T.TBL_NAME='%s'
order by PKV.PART_ID,PKV.INTEGER_IDX;
`

var TABLE_INFO_SQL = `
select S.LOCATION, S.INPUT_FORMAT as type from 
TBLS as T 
join DBS as D on D.DB_ID=T.DB_ID
join SDS as S on T.SD_ID=S.SD_ID
where D.NAME='%s' and T.TBL_NAME='%s';
`

var SHOWTABLES_SQL = `
select TBLS.TBL_NAME 
from TBLS 
join DBS on DBS.DB_ID=TBLS.DB_ID and DBS.NAME='%s';
`
var SHOWSCHEMAS_SQL = `
select DBS.NAME 
from DBS;
`
var SHOWPARTITIONS_SQL = `

`
