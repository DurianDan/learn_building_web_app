# this file is like the .env file, 
# storing many environment variables
# file will be loaded, using app.config.from_file("config.py")

import os

# Get the absolute path of the current directory
basedir = os.path.abspath(os.path.dirname(__file__))

SQLALCHEMY_DATABASE_URI = "mysql://root:mypassword@localhost:3306/Daree"

# Define the SQLALCHEMY_TRACK_MODIFICATIONS variable that tells SQLAlchemy whether to track changes to the database
# We will set it to False to avoid unnecessary overhead
SQLALCHEMY_TRACK_MODIFICATIONS = False
