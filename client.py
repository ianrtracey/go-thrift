import sys
import uuid
from datetime import datetime

from thrift.transport import TSocket, TTransport
from thrift.protocol import TBinaryProtocol
from contact import ContactSvc, ttypes

socket = TSocket.TSocket('localhost', 9090)
transport = TTransport.TFramedTransport(socket)
protocol = TBinaryProtocol.TBinaryProtocol(transport)
client = ContactSvc.Client(protocol)

transport.open()

contact_attributes = ttypes.Contact(uuid.uuid4().hex, 'Ian', '111-1111', 'ian@test.com',
					datetime.now().isoformat())

contact = client.create(contact_attributes)

contacts = client.fetch()
print contacts


