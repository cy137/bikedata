package main

type Station struct {
	Id                 int     `xml:"id"`
	Name               string  `xml:"name"`
	TerminalName       int     `xml:"terminalName"`
	LastCommWithServer string  `xml:"lastCommWithServer"`
	Lat                float32 `xml:"lat"`
	Long               float32 `xml:"long"`
	Installed          bool    `xml:"installed"`
	Locked             bool    `xml:"locked"`
	InstallDate        string  `xml:"installDate"`
	RemovalDate        string  `xml:"removalDate"`
	Temporary          bool    `xml:"temporary"`
	Public             bool    `xml:"public"`
	NbBikes            int     `xml:"nbBikes"`
	NbEmptyDocks       int     `xml:nbEmptyDocks"`
	LatestUpdateTime   string  `xml:"latestUpdateTime"`
}
