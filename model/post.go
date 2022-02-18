package model

//==========================================================
//Post Management
type Post struct {
	Pid         string `json:"pid,omitempty"`
	ChanId      string `json:"chanid,omitempty"`
	ChanName    string `json:"channame,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Title       string `json:"title,omitempty"`
	PcontentRaw string `json:"pcontentraw,omitempty"`
	Pcontent    string `json:"pcontent,omitempty"`
	CaptionUrl  string `json:"captionurl,omitempty"`
	CoverUrl    string `json:"coverurl,omitempty"`
	Ptype       string `json:"ptype,omitempty"`
	SourceUrl   string `json:"sourceurl,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Issuer      string `json:"issuer,omitempty"`
	IssueTime   int64  `json:"issuetime,omitempty"`
	EditTime    int64  `json:"edittime,omitempty"`
	Counts      int    `json:"counts,omitempty"`
	Pro         int    `json:"pro,omitempty"`
	Con         int    `json:"con,omitempty"`
	Tpr         int    `json:"tpr,omitempty"`
	NumComs     int    `json:"numcoms,omitempty"`
}
