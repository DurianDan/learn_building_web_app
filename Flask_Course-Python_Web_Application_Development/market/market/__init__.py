#this is the file that contain all the nescesary imports 
#to run run.py file or anyother python file  that contribute to the flask app
from flask import Flask, render_template
from flask_sqlalchemy import SQLAlchemy 

app = Flask(__name__)
app.config["SQLALCHEMY_DATABASE_URI"] = "sqlite:///market.db"
#the "config" is a dictionary inside the flask "app" class
#URI stands for Unifided Resource Identifier
#Here, we point the database key to an actual file on the computer (will be created, if not existed yet)
app.config["SECRET_KEY"] = "ad40ef1cfa20947d1bec3c56"
#to initialize the database (when not created yet)
#type "python3" in the terminal
'''
    from market import db
*an warning will appear after this*
    db.create_all()
*this command create a db binary file in the project folder*
''' #see the file type_in_terminal.py for furthur information
db = SQLAlchemy(app) #initialize a database for the "app"

from market import routes