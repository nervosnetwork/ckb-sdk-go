package resp

type DBInfo struct {
	Version   string    `json:"version"`
	DB        DBDriver  `json:"db"`
	ConnSize  uint32    `json:"conn_size"`
	CenterId  int64     `json:"center_id"`
	MachineId int64     `json:"machine_id"`
}

type DBDriver string
const (
    PostgreSQL DBDriver = "PostgreSQL"
    MySQL      DBDriver = "MySQL"
    SQLite     DBDriver = "SQLite"
)
