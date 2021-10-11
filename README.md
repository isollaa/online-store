# online-store
We are members of the engineering team of an online store. When we look at ratings for our online store application, we received the following facts: 
1. Customers were able to put items in their cart, check out, and then pay. After several days, many of our customers received calls from our Customer Service department stating that their orders have been canceled due to stock unavailability. 
2. These bad reviews generally come within a week after our 12.12 event, in which we held a large flash sale and set up other major discounts to promote our store. 
After checking in with our Customer Service and Order Processing departments, we received the following additional facts: 
1. Our inventory quantities are often misreported, and some items even go as far as having a negative inventory quantity. 2. The misreported items are those that performed very well on our 12.12 event. 
3. Because of these misreported inventory quantities, the Order Processing department was unable to fulfill a lot of orders, and thus requested help from our Customer Service department to call our customers and notify them that we have had to cancel their orders. 
#


# analysis
Based on the stated facts above, there is a thing that might be the source of issue:

- there is a possibility that users get product availability information in real time because the number of requests that affect product stock occurs very quickly and almost simultaneously
#

# solution
- make sure to lock product on each request, so when a request is made simultaneously, the user will be forced to wait for each other to get the real time data
#

# how to run
- clone this [repo](https://github.com/isollaa/online-store.git)
- open ``config/env.json`` and set based on your environment.  
- go build and run ``./online-store``
- open and run [documenter](https://documenter.getpostman.com/view/10609164/UV5RkzXn#58a94317-4439-4b1a-bfbb-2bd0f8be4ca3) 
- for the first try make sure to login using 
```
username : "admin"  
password : "admin"
```

## unit-test
- open test package
- make sure there is item on table ``item``
- if table ``item`` is empty, create dummt item on test/item package and test ``TestCreateItem()`` function
- move to ``test/cart`` and test ``TestCreateCart()`` function
- done