// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package contact

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

type Contact struct {
	Id      string  `thrift:"id,1,required" json:"id"`
	Name    string  `thrift:"name,2,required" json:"name"`
	Phone   string  `thrift:"phone,3,required" json:"phone"`
	Email   *string `thrift:"email,4" json:"email"`
	Created string  `thrift:"created,5,required" json:"created"`
}

func NewContact() *Contact {
	return &Contact{}
}

func (p *Contact) GetId() string {
	return p.Id
}

func (p *Contact) GetName() string {
	return p.Name
}

func (p *Contact) GetPhone() string {
	return p.Phone
}

var Contact_Email_DEFAULT string

func (p *Contact) GetEmail() string {
	if !p.IsSetEmail() {
		return Contact_Email_DEFAULT
	}
	return *p.Email
}

func (p *Contact) GetCreated() string {
	return p.Created
}
func (p *Contact) IsSetEmail() bool {
	return p.Email != nil
}

func (p *Contact) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *Contact) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.Id = v
	}
	return nil
}

func (p *Contact) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Name = v
	}
	return nil
}

func (p *Contact) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Phone = v
	}
	return nil
}

func (p *Contact) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Email = &v
	}
	return nil
}

func (p *Contact) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.Created = v
	}
	return nil
}

func (p *Contact) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Contact"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *Contact) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("id", thrift.STRING, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:id: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Id)); err != nil {
		return fmt.Errorf("%T.id (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:id: %s", p, err)
	}
	return err
}

func (p *Contact) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:name: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Name)); err != nil {
		return fmt.Errorf("%T.name (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:name: %s", p, err)
	}
	return err
}

func (p *Contact) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("phone", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:phone: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Phone)); err != nil {
		return fmt.Errorf("%T.phone (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:phone: %s", p, err)
	}
	return err
}

func (p *Contact) writeField4(oprot thrift.TProtocol) (err error) {
	if p.IsSetEmail() {
		if err := oprot.WriteFieldBegin("email", thrift.STRING, 4); err != nil {
			return fmt.Errorf("%T write field begin error 4:email: %s", p, err)
		}
		if err := oprot.WriteString(string(*p.Email)); err != nil {
			return fmt.Errorf("%T.email (4) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 4:email: %s", p, err)
		}
	}
	return err
}

func (p *Contact) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("created", thrift.STRING, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:created: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Created)); err != nil {
		return fmt.Errorf("%T.created (5) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:created: %s", p, err)
	}
	return err
}

func (p *Contact) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Contact(%+v)", *p)
}
