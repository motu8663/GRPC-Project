# GRPC-Project
**Prerequisites**
Golang (version 1.16+)
Protocol Buffers compiler (protoc)

**gRPC Service Endpoin**
Fetch User by ID
Request: GetUserByIDRequest { id: int32 }
Response: GetUserByIDResponse { user: User }

Fetch Users by List of IDs
Request: GetUsersByIDsRequest { ids: repeated int32 }
Response: GetUsersByIDsResponse { users: repeated User }

Search Users
Request: SearchUsersRequest { city: string, phone: int64, married: bool }
Response: SearchUsersResponse { users: repeated User }

**Configuration Details**
Ensure that the protoc compiler is installed and accessible to PATH.
