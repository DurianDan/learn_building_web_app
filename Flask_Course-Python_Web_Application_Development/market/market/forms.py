# first, you must install flask-wtf through pip3
# flask-wtf is a package that provide interactive secure user interface

# creating logins forms for the users

from flask_wtf import FlaskForm
from wtforms import StringField, PasswordField, SubmitField
# StringField
# PasswordField
# SubmitField

class RegisterForm(FlaskForm):
    username = StringField(label="Ur mom")
    email_address = StringField(label="Email Address")
    password1 = PasswordField(label="Password")
    password2 = PasswordField(label="Confirm Password")
    #then compare 2 passwords to validate new password created
    submit = SubmitField(label='Create Account')
    # submit changes, save the user's data