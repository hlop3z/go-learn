# Go `encoding`

In Go, the `encoding` package provides several encoders and decoders for different data formats. Besides `encoding/json` for JSON, there are a few other commonly used encoders and decoders. Here's an overview of the main ones:

## Summary of Common Encoders/Decoders

- **`encoding/json`**: JSON format (text-based, lightweight).
- **`encoding/xml`**: XML format (hierarchical, verbose).
- **`encoding/csv`**: CSV format (tabular, text-based).
- **`encoding/gob`**: Gob format (Go-specific, binary).
- **`encoding/base64`**: Base64 (for encoding binary data as text).
- **`encoding/hex`**: Hexadecimal encoding (binary data as hex).
- **`encoding/pem`**: PEM format (used in cryptography).
- **`encoding/asn1`**: ASN.1 format (used in cryptographic protocols).

Each of these encoders provides functionality for encoding data to a specific format and decoding data from that format. Depending on your use case, you can choose the appropriate encoder for tasks such as serialization, communication, or file storage.

### 1. **encoding/json** - JSON Encoding/Decoding

- **Package**: `encoding/json`
- **Purpose**: For encoding (marshaling) and decoding (unmarshaling) Go data structures to and from JSON format.

**Functions**:

- `json.Marshal`: Encodes Go data into JSON format.
- `json.Unmarshal`: Decodes JSON data into Go structures.
- `json.NewEncoder`: Creates a new encoder that can write JSON data to an `io.Writer`.
- `json.NewDecoder`: Creates a new decoder that reads JSON from an `io.Reader`.

### 2. **encoding/xml** - XML Encoding/Decoding

- **Package**: `encoding/xml`
- **Purpose**: For encoding (marshaling) and decoding (unmarshaling) Go data structures to and from XML format.

**Functions**:

- `xml.Marshal`: Encodes Go data into XML format.
- `xml.Unmarshal`: Decodes XML data into Go structures.
- `xml.NewEncoder`: Creates a new encoder that can write XML data to an `io.Writer`.
- `xml.NewDecoder`: Creates a new decoder that reads XML from an `io.Reader`.

### 3. **encoding/csv** - CSV Encoding/Decoding

- **Package**: `encoding/csv`
- **Purpose**: For encoding (marshaling) and decoding (unmarshaling) Go data to and from CSV format.

**Functions**:

- `csv.NewReader`: Reads and parses CSV data from an `io.Reader`.
- `csv.NewWriter`: Writes CSV data to an `io.Writer`.
- `csv.Reader.Read`: Reads a single CSV record.
- `csv.Writer.Write`: Writes a single CSV record.

### 4. **encoding/gob** - Gob Encoding/Decoding

- **Package**: `encoding/gob`
- **Purpose**: For encoding and decoding Go data structures to and from the Gob binary format. Gob is Go's native serialization format and is particularly efficient for Go-specific data types.

**Functions**:

- `gob.NewEncoder`: Creates a new encoder that writes Gob data to an `io.Writer`.
- `gob.NewDecoder`: Creates a new decoder that reads Gob data from an `io.Reader`.
- `gob.Encode`: Encodes a Go value into the Gob format.
- `gob.Decode`: Decodes Gob data into a Go value.

### 5. **encoding/base64** - Base64 Encoding/Decoding

- **Package**: `encoding/base64`
- **Purpose**: For encoding and decoding binary data in Base64 format, which is commonly used for encoding binary data as text (such as for email or HTTP data).

**Functions**:

- `base64.StdEncoding.EncodeToString`: Encodes data into a Base64 string.
- `base64.StdEncoding.DecodeString`: Decodes a Base64 encoded string.
- `base64.NewEncoder`: Encodes data to an `io.Writer` in Base64.
- `base64.NewDecoder`: Decodes Base64 data from an `io.Reader`.

### 6. **encoding/hex** - Hex Encoding/Decoding

- **Package**: `encoding/hex`
- **Purpose**: For encoding and decoding binary data into hexadecimal format.

**Functions**:

- `hex.EncodeToString`: Encodes a byte slice to a hexadecimal string.
- `hex.DecodeString`: Decodes a hexadecimal string into a byte slice.
- `hex.NewEncoder`: Writes hexadecimal-encoded data to an `io.Writer`.
- `hex.NewDecoder`: Reads and decodes hexadecimal data from an `io.Reader`.

### 7. **encoding/pem** - PEM Encoding/Decoding

- **Package**: `encoding/pem`
- **Purpose**: For encoding and decoding data in PEM (Privacy-Enhanced Mail) format, commonly used for certificates and cryptographic keys.

**Functions**:

- `pem.Encode`: Encodes a `pem.Block` to a writer.
- `pem.Decode`: Decodes a PEM-encoded block from a reader.
- `pem.Block`: A structure representing a PEM-encoded block of data, such as a certificate or key.

### 8. **encoding/asn1** - ASN.1 Encoding/Decoding

- **Package**: `encoding/asn1`
- **Purpose**: For encoding and decoding data using the Abstract Syntax Notation One (ASN.1) format, commonly used for certificates, cryptographic keys, and other protocols.

**Functions**:

- `asn1.Marshal`: Encodes Go values to ASN.1 format.
- `asn1.Unmarshal`: Decodes ASN.1 data to Go values.
