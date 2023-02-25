from market import app,db
from flask import render_template, redirect, url_for
from market.models import Item, Users
from market.forms import RegisterForm
@app.route("/") #decorator for the main page 
@app.route("/home") #decorator for the homepage
#many routes can lead to a fundtion !!!
def home_website():
    return render_template("home.html",title = "Home")

@app.route("/market")
def market_page():
    items = Item.query.all()
    #return all the items that have been stored in our database
    return render_template("market.html",items= items)
    #the "item" is a key, can be change like variable names
    #the keys send data to the html file
    #check the market.html file and find the variable "items" to understand

@app.route("/register", methods=['GET','POST'])
# the methods'  variables will let the web
# hadle request data 
# => MUST HAVE
def register_page():
    form = RegisterForm() # from the forms.py file
    if form.validate_on_submit():
        # the user_to_create contains the data that
        # the user has just enter the the FIELDS and
        # click on the SUBMIT button
        user_to_create = Users(username=form.username.data,
                               email_address=form.email_address.data,
                               password_hash=form.password1.data)
        db.session.add(user_to_create)
        db.session.commit()
        # After the user has registered
        # redirect to a new page
        return redirect(url_for('market_page'))
    return render_template("register.html",form=form)