import Service2_pb2 as ServiceTest

obj = ServiceTest.Employees()
obj.Name = 'Kira'
obj.age = 20
print(obj.SerializeToString())

with open('output.bin','wb') as f:
    f.write(obj.SerializeToString())
json_data = '{"Name": "Adnan","Age": 20}'
with open('output.json','w',encoding='utf8') as f:
    f.write(json_data)