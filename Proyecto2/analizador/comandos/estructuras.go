package comandos

type MBR struct {
	Mbr_tamano         int64
	Mbr_fecha_creacion [19]byte
	Mbr_dsk_signature  int64
	Dsk_fit            byte
	Mbr_partition      [4]Particion
}

type Particion struct {
	Part_status byte
	Part_type   byte
	Part_fit    byte
	Part_start  int64
	Part_size   int64
	Part_name   [16]byte
}

type EBR struct {
	Part_status byte
	Part_fit    byte
	Part_start  int64
	Part_size   int64
	Part_next   int64
	Part_name   [16]byte
}
