
struct Contact {
    1: required string id,
    2: required string name,
    3: required string phone,
    4: optional string email,
    5: required string created 
}

service ContactSvc {
    Contact create(1:Contact contact),
    Contact read(1:string contactId),
    Contact update(1:Contact contact),
    void destroy(1:string contactId),
    void reset(),
    list<string> fetch()
}
