from flask import Flask
from db import db 

app = Flask(__name__)
app.config.from_file('config.py')
db.init_app(app)

