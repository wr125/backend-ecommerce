A backend ecommerce app with authentication and authorisation using mongoDB and jwt tokens. Can you docker compose but is setup to use a db connection string in the database file. Just download mongoDB and the DB is called backend. Use postman for parsing raw json input. So for example;

POST(SignUp)
http://localhost:3000/users/signup
{
  "first_name": "Robert",
  "last_name": "Willis",
  "email": "rob@gmail.com",
  "password": "ahgs321y",
  "phone": "+447961867631"
}

Type: go run main.go

POST(Login)
http://localhost:3000/users/login
{
"email": "rob@gmail.com",
"password": "ahgs321y"
}


Admin and Product Function (Post)
http://localhost:3000/admin/addproduct
{
  "product_name": "Alienware x15",
  "price": 2200,
  "rating": 9,
  "image": "alienware.jpg"
}

Search Product via regex (Get)
http://localhost:3000/users/search?name=al
{
    "Product_ID": "818152fa9f29be942bd9df91",
    "product_name": "Alienware x15",
    "price": 2200,
    "rating": 9,
    "image": "1.jpg"
  }

*Adding the Products to the Cart (GET REQUEST)

http://localhost:3000/addtocart?id=xxxproduct_idxxx&userID=xxxxxxuser_idxxxxxx

*Removing Item From the Cart (GET REQUEST)

http://localhost:3000/addtocart?id=xxxxxxx&userID=xxxxxxxxxxxx

*Listing the item in the users cart (GET REQUEST) and total price

http://localhost:3000/listcart?id=xxxxxxuser_idxxxxxxxxxx

*Addding the Address (POST REQUEST)

http://localhost:3000/addadress?id=user_id**\*\***\***\*\***

The Address array is limited to two values home and work address more than two address is not acceptable

*Editing the Home Address(PUT REQUEST)

http://localhost:3000/edithomeaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

*Editing the Work Address(PUT REQUEST)

http://localhost:3000/editworkaddress?id=xxxxxxxxxxuser_idxxxxxxxxxxxxxxx

*Delete Addresses(GET REQUEST)

http://localhost:3000/deleteaddresses?id=xxxxxxxxxuser_idxxxxxxxxxxxxx

delete both addresses

*Cart Checkout Function and placing the order(GET REQUEST)

After placing the order the items have to be deleted from cart functonality added

http://localhost:3000?id=xxuser_idxxx

*Instantly Buying the Products(GET REQUEST) http://localhost:3000?userid=xxuser_idxxx&pid=xxxxproduct_idxxxx