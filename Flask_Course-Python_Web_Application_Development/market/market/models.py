from market import db

#each class or the "Model" inside the database is actually a table
#later on, columns and rows will be added to this Item class (or this table)

class Users(db.Model):
    '''This "Users Model" stores the users'input data like password'''
    id = db.Column(db.Integer(),primary_key=True)
    username = db.Column(db.String(length=30), unique=True, nullable=False)
    # when the user input string data longer than the "length" variable,
    # the string will be truncated 
    email_address = db.Column(db.String(length=50 ), unique=True, nullable=False )
    password_hash = db.Column(db.String(length=60), nullable=False)
    budget = db.Column(db.Integer(), nullable=False, default=1000)
    #the budget column has the "default" value of 1000
    #Because the customer not always have the money beforehand
    #We will give the customer 1000 "coins" at the user signing up

    items = db.relationship('Item',backref='owned_user',lazy=True)
    #this "items" instance links to the "Item" Model
    #to point the "user rows" to the items that they have owned
    '''the "lazy" parameter indicates the steps taken to turn data from binary (in the .db file) into usable data
    If lazy==True or lazy=="select": it will load all relationships to the end
    If lazy==Flase or lazy=="joined: it will load all relationship to the end and also each steps'''
    #this "backref" is the reference to the user, searching from the "Item" Model
    #Eg: when viewing the "Iphone 12" in the Item Model, we can see the backref pointing to the "users",in the "User" Model, that have owned that "Iphone 12"

    def __repf__(self):
        return f"User {self.username}"





class Item(db.Model):
    '''this "Item Model" stores the items table with their prices and descriptions'''
    id = db.Column(db.Integer(),primary_key=True)
    #id column is a standard/must have column in a table
    #primary keys

    name = db.Column(db.String(length=30), nullable=False, unique=True)
    #name is a column that contains "string" data
    #the strings in name column has "max length" of 30
    #name column canot be "NUll" or "None", cannot have duplicates 
    price = db.Column(db.Integer(),nullable=False)
    barcode = db.Column(db.String(length=12),nullable=False,unique=True)
    description = db.Column(db.String(length=1024),nullable=False,unique=True)

    owner = db.Column(db.Integer(),db.ForeignKey('users.id'))

    def __repr__(self):
        #what does this mean ? ;)
        #compare this to __str__
        return f"Item {self.name}"
