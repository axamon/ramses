package main

import "time"

// Jerk contiene le informazioni sulle cadute di sessioni PPP
// che si sono verificate.
type Jerk struct {
	NasName   string
	Timestamp time.Time
	pppValue  float64
}

// Jerks è un insieme di Jerk
type Jerks []Jerk

// Configuration tiene gli elementi di configurazione
type Configuration struct {
	IPDOMUser          string
	IPDOMPassword      string
	IPDOMSnmpReceiver  string
	IPDOMSnmpPort      uint16
	IPDOMSnmpCommunity string
	Sigma              float64
	Soglia             float64
	NasInventory       string
	NasDaIgnorare      string
	URLSessioniPPP     string
	URLTail7d          string
	SmtpPort           int
	SmtpServer         string
	SmtpUser           string
	SmtpPassword       string
	SmtpSender         string
	SmtpFrom           string
	SmtpTo             string
}

// TNAS is a structure Type for contain NAS interface
type TNAS struct {
	ID                int
	Name              string
	Description       string
	Location          string
	Service           string
	ServiceAdded      string
	Pop               string
	ManIPAddress      string
	Network           string
	Centrale          string
	Alias             string
	Domain            string
	CreateTime        string
	ChangeTime        string
	SysContact        string
	SysDescr          string
	SysLocation       string
	SysName           string
	SysObject         string
	Aisle             string
	Altitude          string
	ChassisUUID       string
	FRUNumber         string
	FRUSerialNumber   string
	FWRevision        string
	HWRevision        string
	Manufacturer      string
	ManufacturingDate string
	Model             string
	ChassisName       string
	PartNumber        string
	RfidTag           string
	RackPosition      string
	RelativePosition  string
	CdmRow            string
	SwRevision        string
	SerialNumber      string
	SystemBoardUUID   string
	CdmType           string
	UserTracking      string
	XCoordinate       string
	YCoordinate       string
	PhysicalIndex     string
	VendorType        string
	ClassName         string
	Uptime            string
	Services          string
	InterfaceCount    string
	IsIPForwarding    string
	AccessIPAddress   string
	AccessProtocol    string
	DiscoveryTime     string
	Timestamp         string
	Routing           string
	Metrics           string
	Events            string
	Device            string
	Interfaces        string
	Fans              string
	Cards             string
	Slots             string
	Psus              string
	Configurations    string
}
