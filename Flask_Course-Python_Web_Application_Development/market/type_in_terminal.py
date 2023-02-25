from market import db
db.create_all()
'''
# to initialize or call the database (when not created yet)
# type "python3" in the terminal
>>from market import db
#an warning will appear after this
>>db.create_all()
# this command create a db binary file in the project folder
'''
# to add instances to the table/model "Item"
from market import Item
item1 = Item(name="Iphone 10",price=300,barcode="2039809qvj",description="test-duh")
db.session.add(item1)
db.session.commit()
db.drop_all() # delete all data in the database
db.create_all() # call the database
db.session.rollback() # return to the start of the database, erase all the data query before

# generate a random number, uses as a secrete key
# for database
import os
print(os.urandom(12).hex())