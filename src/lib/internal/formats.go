package internal

type ArchiveFormat struct {
    Ext string
}

var ArchiveFormats = map[string]ArchiveFormat{
    "7zip": {Ext: "7z"},
    "zip":  {Ext: "zip"},
    "zst":  {Ext: "zst"},
    "tar":  {Ext: "tar"}
}

var ArchiveProfiles = map[string]struct{}{
    "fast":    {},
    "normal":  {},
    "extreme": {},
    "store":   {}
}
