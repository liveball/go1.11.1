// Code generated by go tool dist; DO NOT EDIT.
// This is a bootstrap copy of /Users/fpf/Downloads/go1.11.1/src/cmd/link/internal/sym/symkind_string.go

//line /Users/fpf/Downloads/go1.11.1/src/cmd/link/internal/sym/symkind_string.go:1
// Code generated by "stringer -type=SymKind"; DO NOT EDIT.

package sym

import "strconv"

const _SymKind_name = "SxxxSTEXTSELFRXSECTSTYPESSTRINGSGOSTRINGSGOFUNCSGCBITSSRODATASFUNCTABSELFROSECTSMACHOPLTSTYPERELROSSTRINGRELROSGOSTRINGRELROSGOFUNCRELROSGCBITSRELROSRODATARELROSFUNCTABRELROSTYPELINKSITABLINKSSYMTABSPCLNTABSELFSECTSMACHOSMACHOGOTSWINDOWSSELFGOTSNOPTRDATASINITARRSDATASBSSSNOPTRBSSSTLSBSSSXREFSMACHOSYMSTRSMACHOSYMTABSMACHOINDIRECTPLTSMACHOINDIRECTGOTSFILEPATHSCONSTSDYNIMPORTSHOSTOBJSDWARFSECTSDWARFINFOSDWARFRANGESDWARFLOCSDWARFMISC"

var _SymKind_index = [...]uint16{0, 4, 9, 19, 24, 31, 40, 47, 54, 61, 69, 79, 88, 98, 110, 124, 136, 148, 160, 173, 182, 191, 198, 206, 214, 220, 229, 237, 244, 254, 262, 267, 271, 280, 287, 292, 304, 316, 333, 350, 359, 365, 375, 383, 393, 403, 414, 423, 433}

func (i SymKind) String() string {
	if i >= SymKind(len(_SymKind_index)-1) {
		return "SymKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SymKind_name[_SymKind_index[i]:_SymKind_index[i+1]]
}
