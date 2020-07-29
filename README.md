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
<br></br>
> ##### POST {{host}}/magazine/add_magazine_bullet
> add magazine bullet
```$xslt
{
    "id": "bb466b50-d131-11ea-bb01-309c23fed129",
    "qty": 3
}
```
<br></br>
> ##### GET {{host}}/magazine/attach_magazine
> attach magazine to gun
```$xslt
id = ca1fad7a-d131-11ea-bb01-309c23fed129
```
<br></br>
> ##### GET {{host}}/magazine/detach_magazine
> detach magazine from gun
```$xslt
id = ca1fad7a-d131-11ea-bb01-309c23fed129
```
<br></br>
> ##### GET {{host}}/magazine/verify
> verify magazine

<br></br>
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
<br></br>
> ##### POST {{host}}/store/add_product_quantity
> add new product quantity
```$xslt
{
    "id": "d2c858ab-d159-11ea-b549-309c23fed129",
    "qty": 5
}
```
<br></br>
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
<br></br>
> ##### GET {{host}}/store/verify_order
> add new order
```$xslt
id = ae8901c1-d166-11ea-b87a-309c23fed129
```

 <br></br>
 > by *Ahmad Reza Musthafa*