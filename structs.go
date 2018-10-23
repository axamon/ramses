package main

type TDATA []struct {
	Time  float64 `json:"time"`
	Value float64 `json:"value"`
}

type XrsMi001Stru struct {
	NetThroughputOut struct {
		XrsMi001 struct {
			A110100EthernetTX struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"A/1-10/100-Ethernet-TX"`
			B110100EthernetTX struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"B/1-10/100-Ethernet-TX"`
			Lag99LAGGroup00000000LOGICORMI595Ae5OFFRAMPToRMI595200G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-99-LAG-Group-000000/00-LOGICO-R-MI595-ae5-OFFRAMP-to-R-MI595-200G-"`
			Three2210GigEthernetICRA0340557LocaleCMI284TE131CHUPLINKSEC23 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/2/2-10-Gig-Ethernet-ICR-A03405/57-Locale-C-MI284-TE13/1-CH-UPLINK-SEC-23-"`
			Three2110GigEthernetICRA0340556LocaleCMI284TE103LOCTRANSITMIMI struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/2/1-10-Gig-Ethernet-ICR-A03405/56-Locale-C-MI284-TE10/3-LOC-TRANSIT-MI-MI-"`
			Four1110GigEthernetICRA0340558LocaleCMI285TE103LOCTRANSITMIMI struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"4/1/1-10-Gig-Ethernet-ICR-A03405/58-Locale-C-MI285-TE10/3-LOC-TRANSIT-MI-MI-"`
			Two24100GigEthernetICRA0338453GeograficoXRSRM003113INNERNGCNMIRMLAG3 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/2/4-100-Gig-Ethernet-ICR-A03384/53-Geografico-XRS-RM003-1/1/3-INNER-NGCN-MI-RM-LAG3-"`
			Two23100GigEthernetICRA0335685GeograficoXRSRM003112INNERNGCNMIRMLAG3 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/2/3-100-Gig-Ethernet-ICR-A03356/85-Geografico-XRS-RM003-1/1/2-INNER-NGCN-MI-RM-LAG3-"`
			One23100GigEthernetICRA0335683GeograficoXRSRM001222INNERNGCNMIRMLAG2 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/2/3-100-Gig-Ethernet-ICR-A03356/83-Geografico-XRS-RM001-2/2/2-INNER-NGCN-MI-RM-LAG2-"`
			One14100GigEthernetICRA0335673GeograficoXRSBO001111PRINGCNMIBOLAG34 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/1/4-100-Gig-Ethernet-ICR-A03356/73-Geografico-XRS-BO001-1/1/1-PRI-NGCN-MI-BO-LAG34-"`
			Two14100GigEthernetICRA0335675GeograficoXRSVR001111PRINGCNMIVRLAG33 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/1/4-100-Gig-Ethernet-ICR-A03356/75-Geografico-XRS-VR001-1/1/1-PRI-NGCN-MI-VR-LAG33-"`
			Three14100GigEthernetICRA0338509LOCALEXRSMI009112PRINGCNMIMIMRTG100GLAG32 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/1/4-100-Gig-Ethernet-ICR-A03385/09-LOCALE-XRS-MI009-1/1/2-PRI-NGCN-MI-MI-MRTG-100G-LAG32-"`
			One21100GigEthernetICRA0335678LOCALEXRSMI009111PRINGCNMIMIMRTG100GLAG32 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/2/1-100-Gig-Ethernet-ICR-A03356/78-LOCALE-XRS-MI009-1/1/1-PRI-NGCN-MI-MI-MRTG-100G-LAG32-"`
			Lag101LAGGroupLAG101ToNETFLIX struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-101-LAG-Group-LAG-101-to-NETFLIX-"`
			Six13100GigEthernetICRC0024760MetropolitanoMISEEWEBCalderaNETFLIXMRTG100G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/1/3-100-Gig-Ethernet-ICR-C00247/60-Metropolitano-MI-SEEWEB-Caldera-NETFLIX-MRTG-100G-"`
			One11100GigEthernetICRC0022758MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/1/1-100-Gig-Ethernet-ICR-C00227/58-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			Two13100GigEthernetICRC0022806MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/1/3-100-Gig-Ethernet-ICR-C00228/06-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			One12100GigEthernetICRC0022759MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/1/2-100-Gig-Ethernet-ICR-C00227/59-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			Three13100GigEthernetICRA0335681MetropolitanoXRSMI003313INNERNGCNMIMILAG1 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/1/3-100-Gig-Ethernet-ICR-A03356/81-Metropolitano-XRS-MI003-3/1/3-INNER-NGCN-MI-MI-LAG1-"`
			One13100GigEthernetICRC0022760MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/1/3-100-Gig-Ethernet-ICR-C00227/60-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			Two21100GigEthernetICRA0335669LOCALEXRSMI007111PRINGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/2/1-100-Gig-Ethernet-ICR-A03356/69-LOCALE-XRS-MI007-1/1/1-PRI-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Two22100GigEthernetICRA0335684GeograficoXRSRM003111INNERNGCNMIRMLAG3 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/2/2-100-Gig-Ethernet-ICR-A03356/84-Geografico-XRS-RM003-1/1/1-INNER-NGCN-MI-RM-LAG3-"`
			Three11100GigEthernetICRA0335679MetropolitanoXRSMI003311INNERNGCNMIMILAG1 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/1/1-100-Gig-Ethernet-ICR-A03356/79-Metropolitano-XRS-MI003-3/1/1-INNER-NGCN-MI-MI-LAG1-"`
			Three12100GigEthernetICRA0335680MetropolitanoXRSMI003312INNERNGCNMIMILAG1 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"3/1/2-100-Gig-Ethernet-ICR-A03356/80-Metropolitano-XRS-MI003-3/1/2-INNER-NGCN-MI-MI-LAG1-"`
			Four21100GigEthernetICRA0340805LOCALEXRSMI007112PRINGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"4/2/1-100-Gig-Ethernet-ICR-A03408/05-LOCALE-XRS-MI007-1/1/2-PRI-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Two11100GigEthernetICRC0022804MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/1/1-100-Gig-Ethernet-ICR-C00228/04-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			Two12100GigEthernetICRC0022805MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"2/1/2-100-Gig-Ethernet-ICR-C00228/05-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG100-"`
			Four23100GigEthernetICRA0345634LocaleRMI595Et200OFFRAMPMRTG100GLAG99 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"4/2/3-100-Gig-Ethernet-ICR-A03456/34-Locale-R-MI595-et-2/0/0-OFFRAMP-MRTG-100G-LAG99-"`
			Five22100GigEthernetICRA0345635LocaleRMI595Et900OFFRAMPMRTG100GLAG99 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/2/2-100-Gig-Ethernet-ICR-A03456/35-Locale-R-MI595-et-9/0/0-OFFRAMP-MRTG-100G-LAG99-"`
			Four22100GigEthernetICRA0342419MetropolitanoXRSMI003422INNERNGCNMIMILAG1 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"4/2/2-100-Gig-Ethernet-ICR-A03424/19-Metropolitano-XRS-MI003-4/2/2-INNER-NGCN-MI-MI-LAG1-"`
			Five13100GigEthernetICRC0023877MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/1/3-100-Gig-Ethernet-ICR-C00238/77-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG-100-"`
			Five11100GigEthernetICRC0023878MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/1/1-100-Gig-Ethernet-ICR-C00238/78-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG-100-"`
			Five12100GigEthernetICRC0023879MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/1/2-100-Gig-Ethernet-ICR-C00238/79-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG-100-"`
			Five14100GigEthernetICRC0023875MetropolitanoHNE500CMICLD50SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/1/4-100-Gig-Ethernet-ICR-C00238/75-Metropolitano-HNE500C-MI-CLD50-SEABONE-LAG-100-"`
			Five24100GigEthernetICRA0347575GeograficoXRSRM003322INNERNGCNMIRMLAG3 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/2/4-100-Gig-Ethernet-ICR-A03475/75-Geografico-XRS-RM003-3/2/2-INNER-NGCN-MI-RM-LAG3-"`
			Six12100GigEthernetICRA0347572GeograficoXRSRM001322INNERNGCNRMMILAG2 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/1/2-100-Gig-Ethernet-ICR-A03475/72-Geografico-XRS-RM001-3/2/2-INNER-NGCN-RM-MI-LAG2-"`
			Five23100GigEthernetICRA0347201GeograficoXRSBO001112PRINGCNMIBOLAG34100G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/2/3-100-Gig-Ethernet-ICR-A03472/01-Geografico-XRS-BO001-1/1/2-PRI-NGCN-MI-BO-LAG34-100G-"`
			Four24100GigEthernetICRA0347183GeograficoXRSRM003321INNERNGCNRMMILAG3 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"4/2/4-100-Gig-Ethernet-ICR-A03471/83-Geografico-XRS-RM003-3/2/1-INNER-NGCN-RM-MI-LAG3-"`
			Six11100GigEthernetICRA0347182GeograficoXRSRM001321INNERNGCNRMMILAG2 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/1/1-100-Gig-Ethernet-ICR-A03471/82-Geografico-XRS-RM001-3/2/1-INNER-NGCN-RM-MI-LAG2-"`
			Lag34LAGGroup00000000LOGICOXRSBO001LAG34PRINGCNMIBOToXRSBO001300G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-34-LAG-Group-000000/00-LOGICO-XRS-BO001-LAG34-PRI-NGCN-MI-BO-to-XRS-BO001-300G-"`
			Seven22100GigEthernetICRA0349223MetropolitanoXRSMI003722INNERNGCNMIMILAG1 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/2/2-100-Gig-Ethernet-ICR-A03492/23-Metropolitano-XRS-MI003-7/2/2-INNER-NGCN-MI-MI-LAG1-"`
			Lag2LAGGroup00000000LOGICOXRSRM001LAG2INNERNGCNMIRMToXRSRM001600G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-2-LAG-Group-000000/00-LOGICO-XRS-RM001-LAG2-INNER-NGCN-MI-RM-to-XRS-RM001-600G-"`
			Lag33LAGGroup00000000LOGICOXRSVR001LAG33PRINGCNVRMIToXRSVR001300G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-33-LAG-Group-000000/00-LOGICO-XRS-VR001-LAG33-PRI-NGCN-VR-MI-to-XRS-VR001-300G-"`
			Lag3LAGGroup00000000LOGICOXRSRM003LAG3INNERNGCNMIRMToXRSRM003600G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-3-LAG-Group-000000/00-LOGICO-XRS-RM003-LAG3-INNER-NGCN-MI-RM-to-XRS-RM003-600G-"`
			Seven12100GigEthernetA0349234GEOGRAFICOXRSRM003424PRINGCNMIRM100G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/1/2-100-Gig-Ethernet-A03492/34-GEOGRAFICO-XRS-RM003-4/2/4-PRI-NGCN-MI-RM-100G-"`
			Lag1LAGGroup00000000LOGICOXRSMI003LAG1INNERNGCNMIMIToXRSMI003500G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-1-LAG-Group-000000/00-LOGICO-XRS-MI003-LAG1-INNER-NGCN-MI-MI-to-XRS-MI003-500G-"`
			One24100GigEthernetICRA0338451GeograficoXRSRM001223INNERNGCNMIRMLAG2 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"1/2/4-100-Gig-Ethernet-ICR-A03384/51-Geografico-XRS-RM001-2/2/3-INNER-NGCN-MI-RM-LAG2-"`
			Seven11100GigEthernetA0349260GEOGRAFICOXRSBO001114PRINGCNMIBO100G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/1/1-100-Gig-Ethernet-A03492/60-GEOGRAFICO-XRS-BO001-1/1/4-PRI-NGCN-MI-BO-100G-"`
			One014100GigEthernetINATTIVAZIONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/1/4-100-Gig-Ethernet-IN-ATTIVAZIONE-"`
			Five21100GigEthernetICRC0023876MetropolitanoHNE500CMICLD50Hu1204SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"5/2/1-100-Gig-Ethernet-ICR-C00238/76-Metropolitano-HNE500C-MI-CLD50-Hu12/0/4-SEABONE-LAG-100-"`
			One012100GigEthernetfloat64EST struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/1/2-100-Gig-Ethernet-IN-TEST-"`
			Nine23100GigEthernetICRC0023626MetropolitanoHNE500CMICLD50Hu1402SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"9/2/3-100-Gig-Ethernet-ICR-C00236/26-Metropolitano-HNE500C-MI-CLD50-Hu14/0/2-SEABONE-LAG100-"`
			Nine24100GigEthernetICRC0023627MetropolitanoHNE500CMICLD50Hu1300SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"9/2/4-100-Gig-Ethernet-ICR-C00236/27-Metropolitano-HNE500C-MI-CLD50-Hu13/0/0-SEABONE-LAG100-"`
			One011100GigEthernetfloat64EST struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/1/1-100-Gig-Ethernet-IN-TEST-"`
			One022100GigEthernetICRC0023629MetropolitanoHNE500CMICLD50Hu1202SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/2/2-100-Gig-Ethernet-ICR-C00236/29-Metropolitano-HNE500C-MI-CLD50-Hu12/0/2-SEABONE-LAG100-"`
			One021100GigEthernetICRC0023628MetropolitanoHNE500CMICLD50Hu1200SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/2/1-100-Gig-Ethernet-ICR-C00236/28-Metropolitano-HNE500C-MI-CLD50-Hu12/0/0-SEABONE-LAG100-"`
			Seven24100GigEthernetICRA0349229GeograficoXRSRM001424INNERNGCNMIRMLAG2 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/2/4-100-Gig-Ethernet-ICR-A03492/29-Geografico-XRS-RM001-4/2/4-INNER-NGCN-MI-RM-LAG2-"`
			Seven21100GigEthernetICRA0349256GeograficoXRSVR001113PRINGCNMIVR100G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/2/1-100-Gig-Ethernet-ICR-A03492/56-Geografico-XRS-VR001-1/1/3-PRI-NGCN-MI-VR-100G-"`
			Seven13100GigEthernetICRA0349238GeograficoXRSMI009113PRINGCNMIMIMRTG100GECC31072018LAG32 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/1/3-100-Gig-Ethernet-ICR-A03492/38-Geografico-XRS-MI009-1/1/3-PRI-NGCN-MI-MI-MRTG-100G-ECC-31/07/2018-LAG32-"`
			Lag32LAGGroupICR00000000LOGICOXRSMI009LAG32PRINGCNMIMIMRTG300G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-32-LAG-Group-ICR-000000/00-LOGICO-XRS-MI009-LAG32-PRI-NGCN-MI-MI-MRTG-300G-"`
			Seven23100GigEthernetICRC0024930MetropolitanoHNE500CMICLD50Hu0801SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/2/3-100-Gig-Ethernet-ICR-C00249/30-Metropolitano-HNE500C-MI-CLD50-Hu0/8/0/1-SEABONE-LAG100-"`
			Six24100GigEthernetICRC0024929MetropolitanoHNE500CMICLD50Hu01101SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/2/4-100-Gig-Ethernet-ICR-C00249/29-Metropolitano-HNE500C-MI-CLD50-Hu0/11/0/1-SEABONE-LAG100-"`
			Six23100GigEthernetICRC0024928MetropolitanoHNE500CMICLD50Hu0800SEABONELAG100 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/2/3-100-Gig-Ethernet-ICR-C00249/28-Metropolitano-HNE500C-MI-CLD50-Hu0/8/0/0-SEABONE-LAG100-"`
			Six14100GigEthernetICRA0350559LOCALEXRSMI007113PRINGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"6/1/4-100-Gig-Ethernet-ICR-A03505/59-LOCALE-XRS-MI007-1/1/3-PRI-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Eight12100GigEthernetICRC0024933MetropolitanoHNE500CMICLD50Hu01107SEABONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/1/2-100-Gig-Ethernet-ICR-C00249/33-Metropolitano-HNE500C-MI-CLD50-Hu0/11/0/7-SEABONE-"`
			Eight24100GigEthernetICRC0024931MetropolitanoHNE500CMICLD50Hu0807SEABONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/2/4-100-Gig-Ethernet-ICR-C00249/31-Metropolitano-HNE500C-MI-CLD50-Hu0/8/0/7-SEABONE-"`
			Eight11100GigEthernetICRC0024932MetropolitanoHNE500CMICLD50Hu08014SEABONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/1/1-100-Gig-Ethernet-ICR-C00249/32-Metropolitano-HNE500C-MI-CLD50-Hu0/8/0/14-SEABONE-"`
			Eight23100GigEthernetICRC0024934MetropolitanoHNE500CMICLD50Hu011014SEABONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/2/3-100-Gig-Ethernet-ICR-C00249/34-Metropolitano-HNE500C-MI-CLD50-Hu0/11/0/14-SEABONE-"`
			Nine21100GigEthernetINATTIVAZIONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"9/2/1-100-Gig-Ethernet-IN-ATTIVAZIONE-"`
			Lag100LAGGroupICR00000000LOGICOHNE500CMICLD50LAG100SEABONE2200G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-100-LAG-Group-ICR-000000/00-LOGICO-HNE500C-MI-CLD50-LAG100-SEABONE-2200G-"`
			Seven14100GigEthernetA0349255GEOGRAFICOXRSVR001112PRINGCNMIVR100GLAG33 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"7/1/4-100-Gig-Ethernet-A03492/55-GEOGRAFICO-XRS-VR001-1/1/2-PRI-NGCN-MI-VR-100G-LAG33-"`
			Lag30LAGGroupICR00000000LOGICOXRSMI007LAG30PCHNGCNMIMIMRTG700G struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"lag-30-LAG-Group-ICR-000000/00-LOGICO-XRS-MI007-LAG30-P-CH-NGCN-MI-MI-MRTG-700G-"`
			Eight14100GigEthernetICRA0353939LOCALEXRSMI007412PCHNGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/1/4-100-Gig-Ethernet-ICR-A03539/39-LOCALE-XRS-MI007-4/1/2-P-CH-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Eight13100GigEthernetICRA0353938LOCALEXRSMI007411PCHNGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"8/1/3-100-Gig-Ethernet-ICR-A03539/38-LOCALE-XRS-MI007-4/1/1-P-CH-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Nine12100GigEthernetICRA0353941LOCALEXRSMI007414PCHNGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"9/1/2-100-Gig-Ethernet-ICR-A03539/41-LOCALE-XRS-MI007-4/1/4-P-CH-NGCN-MI-MI-MRTG-100G-LAG30-"`
			Nine11100GigEthernetICRA0353940LOCALEXRSMI007413PCHNGCNMIMIMRTG100GLAG30 struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"9/1/1-100-Gig-Ethernet-ICR-A03539/40-LOCALE-XRS-MI007-4/1/3-P-CH-NGCN-MI-MI-MRTG-100G-LAG30-"`
			One013100GigEthernetINATTIVAZIONE struct {
				Data []struct {
					Time  float64 `json:"time"`
					Value float64 `json:"value"`
				} `json:"data"`
			} `json:"10/1/3-100-Gig-Ethernet-IN-ATTIVAZIONE-"`
		} `json:"xrs-mi001"`
	} `json:"net.throughput.out"`
}
