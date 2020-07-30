# Assessment

## Solution 1
### Verify magazine gun
> ##### POST {{host}}/magazine/add_magazine
> create a new magazine
```$xslt
{
    "name": "Magazine B",
    "qty": 12
}
```
> ##### POST {{host}}/magazine/add_magazine_bullet
> add magazine bullet
```$xslt
{
    "id": "bb466b50-d131-11ea-bb01-309c23fed129",
    "qty": 3
}
```
> ##### GET {{host}}/magazine/attach_magazine
> attach magazine to gun
```$xslt
id = ca1fad7a-d131-11ea-bb01-309c23fed129
```
> ##### GET {{host}}/magazine/detach_magazine
> detach magazine from gun
```$xslt
id = ca1fad7a-d131-11ea-bb01-309c23fed129
```
> ##### GET {{host}}/magazine/verify
> verify magazine

> ##### GET {{host}}/magazine/shot
> shot bullets from gun
```$xslt
qty = 1
```

<br></br>
## Solution 2
### Handle high traffic order
> ##### POST {{host}}/store/add_product
> add new product
```$xslt
{
    "code": "ABC1",
    "name": "Sunlight",
    "qty": 20
}
```
> ##### POST {{host}}/store/add_product_quantity
> add new product quantity
```$xslt
{
    "id": "d2c858ab-d159-11ea-b549-309c23fed129",
    "qty": 5
}
```
> ##### POST {{host}}/store/add_order
> add new order
```$xslt
{
	"products":[
		{
			"product_id":"d2c858ab-d159-11ea-b549-309c23fed129",
			"qty":3
		}
	]
}
```
> ##### GET {{host}}/store/verify_order
> add new order
```$xslt
id = ae8901c1-d166-11ea-b87a-309c23fed129
```

<br></br>
## Solution 3
### Solve probability of key

> RULE
> 1. Go to north Y step, then
> 2. Go to east Y step, then
> 3. Go to south Y step

```$xslt
Solution 1
[# # # # # # # #]
[# . . . . . . #]
[# . # # # . . #]
[# X X X # . # #]
[# X # X . . . #]
[# # # # # # # #]

Solution 2
[# # # # # # # #]
[# . . . . X X #]
[# . # # # X X #]
[# . . . # X # #]
[# . # . . X . #]
[# # # # # # # #]

Solution 3
[# # # # # # # #]
[# X X X X X . #]
[# X # # # X . #]
[# X . . # X # #]
[# . # . . X . #]
[# # # # # # # #]

Solution 4
[# # # # # # # #]
[# . . . . X X #]
[# . # # # X X #]
[# . . . # X # #]
[# . # . . . . #]
[# # # # # # # #]

Solution 5
[# # # # # # # #]
[# X X X X X . #]
[# X # # # X . #]
[# . . . # X # #]
[# . # . . X . #]
[# # # # # # # #]

Solution 6
[# # # # # # # #]
[# . . . . X X #]
[# . # # # X X #]
[# . . . # . # #]
[# . # . . . . #]
[# # # # # # # #]


Merged Solution
[# # # # # # # #]
[# X X X X X X #]
[# X # # # X X #]
[# X X X # X # #]
[# X # X . X . #]
[# # # # # # # #]

Total for probability of key = 16 point
```

> RULE
> 1. Go to north Y step, then
> 2. Go to west Y step, then
> 3. Go to south Y step
```

Solution 1
[# # # # # # # #]
[# . . . . . . #]
[# . # # # . . #]
[# X X X # . # #]
[# X # X . . . #]
[# # # # # # # #]

Solution 2
[# # # # # # # #]
[# X X X X X . #]
[# X # # # X . #]
[# X . . # X # #]
[# X # . . X . #]
[# # # # # # # #]

Solution 3
[# # # # # # # #]
[# X X X X X . #]
[# X # # # X . #]
[# X . . # X # #]
[# X # . . . . #]
[# # # # # # # #]

Solution 4
[# # # # # # # #]
[# X X X X X . #]
[# X # # # X . #]
[# X . . # . # #]
[# X # . . . . #]
[# # # # # # # #]

Solution 5
[# # # # # # # #]
[# . . . . X X #]
[# . # # # X X #]
[# . . . # X # #]
[# . # . . X . #]
[# # # # # # # #]


Merged Solution
[# # # # # # # #]
[# X X X X X X #]
[# X # # # X X #]
[# X X X # X # #]
[# X # X . X . #]
[# # # # # # # #]

Total for probability of key = 16 point
```
 <br></br>
 > by *Ahmad Reza Musthafa*