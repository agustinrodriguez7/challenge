# Challenge made for vidflex company

### I did this thinking this is just a challenge:
* I didn't do a high percentage of coverage with unit testing, just the handlers layer.
* I could use a framework for the database, but so that don't take much time, I use builtin library of Go with native queries.
* I did some configurations that should be improved for production environments (for example: cors to dont allow all origins).
* This is my first time using docker without any knowledge before, so it can contain some "bad practices" or mistakes.
* Sensible info like usernames, and passwords can be consumed from env vars to make it scalable, portable and more secure.
*  The algorithm of products insertion is improvable, adding rows named `quantity` to a `products_carts` and `products_orders` table.

### How to run it?
* Go to the folder `infra` run the file called `start_infra.sh` and it will create everything.
* Only in the first time, go the folder `infra/setup` run the file called `db_init.sh`, then copy in the console the lines of the file `db_initialization.sql`, press `enter` and then write `exit`.


#### To test it, there is postman collection inside of the folder `docs` 