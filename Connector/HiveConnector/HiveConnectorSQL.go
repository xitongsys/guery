package HiveConnector

var MD_SQL = `
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

order by sortkey;
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
select SDS.LOCATION as location, PKV.PART_KEY_VAL as pkey from
TBLS as T 
join PARTITIONS as P on T.TBL_ID=P.TBL_ID
join PARTITION_KEY_VALS as PKV on P.PART_ID=PKV.PART_ID
join SDS as SDS on SDS.SD_ID=P.SD_ID
join DBS as D on D.DB_ID=T.DB_ID
where D.NAME='%s' and T.TBL_NAME='%s'
order by PKV.PART_ID,PKV.INTEGER_IDX;
`
