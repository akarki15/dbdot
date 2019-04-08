# dbdot

`dbdot` is a command line tool that generates [DOT](https://en.wikipedia.org/wiki/DOT_(graph_description_language)) description from postgres database schema.

`dbdot` is compiled to platform specific binary, so all you need is the right binary for your machine. Grab the appropriate binary for your architecture from the [latest release here](https://github.com/akarki15/dbdot/releases).

# Demos
#### Demo 1: Produce DOT for all the tables in postgres db `pgguide` and user `kewluser`
```
$ ./dbdot -dbname=pgguide -user=kewluser                                                                                                       [16:33:31]
digraph  {

	node[label=<<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0"><TR><TD colspan="2">products</TD></TR>[<TR><TD>id</TD><TD>integer</TD></TR> <TR><TD>title</TD><TD>character varying</TD></TR> <TR><TD>price</TD><TD>numeric</TD></TR> <TR><TD>created_at</TD><TD>timestamp with time zone</TD></TR> <TR><TD>deleted_at</TD><TD>timestamp with time zone</TD></TR> <TR><TD>tags</TD><TD>ARRAY</TD></TR>]</TABLE>>,shape=plaintext] n1;
	node[label=<<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0"><TR><TD colspan="2">purchase_items</TD></TR>[<TR><TD>id</TD><TD>integer</TD></TR> <TR><TD>purchase_id</TD><TD>integer</TD></TR> <TR><TD>product_id</TD><TD>integer</TD></TR> <TR><TD>price</TD><TD>numeric</TD></TR> <TR><TD>quantity</TD><TD>integer</TD></TR> <TR><TD>state</TD><TD>character varying</TD></TR>]</TABLE>>,shape=plaintext] n2;
	node[label=<<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0"><TR><TD colspan="2">users</TD></TR>[<TR><TD>id</TD><TD>integer</TD></TR> <TR><TD>email</TD><TD>character varying</TD></TR> <TR><TD>password</TD><TD>character varying</TD></TR> <TR><TD>details</TD><TD>USER-DEFINED</TD></TR> <TR><TD>created_at</TD><TD>timestamp with time zone</TD></TR> <TR><TD>deleted_at</TD><TD>timestamp with time zone</TD></TR>]</TABLE>>,shape=plaintext] n3;
	node[label=<<TABLE BORDER="0" CELLBORDER="1" CELLSPACING="0"><TR><TD colspan="2">purchases</TD></TR>[<TR><TD>id</TD><TD>integer</TD></TR> <TR><TD>created_at</TD><TD>timestamp with time zone</TD></TR> <TR><TD>name</TD><TD>character varying</TD></TR> <TR><TD>address</TD><TD>character varying</TD></TR> <TR><TD>state</TD><TD>character varying</TD></TR> <TR><TD>zipcode</TD><TD>integer</TD></TR> <TR><TD>user_id</TD><TD>integer</TD></TR>]</TABLE>>,shape=plaintext] n4;
	n2->n1;
	n2->n4;
	n4->n3;

}
```
#### Demo 2: Produce a schema diagram for all the tables in `pgguide` db for `kewluser`
```
$ ./dbdot -dbname=pgguide -user=kewluser > test.dot && dot -Tpng test.dot -o outfile.png && open outfile.png
```
Above command pipes the DOT description into test.dot and invokes `dot` cli tool to trasform test.dot to a outfile.png. Here's what outfile.png looks like:

![outfile.png](https://raw.githubusercontent.com/akarki15/dbdot/master/images/outfile.png)

#### Demo 3: Whitelist tables and produce schema diagram for them
```
./dbdot -dbname=pgguide -user=kewluser > test.dot --whitelist=purchase_items,purchases && dot -Tpng test.dot -o outfile-whitelisted.png && open outfile-whitelisted.png
```
Here's what outfile-whitelisted.png looks like:

![outfile-whitelisted.png](https://raw.githubusercontent.com/akarki15/dbdot/master/images/outfile-whitelisted.png)

# Flags
dbdot currently supports the following flags:
```
  -whitelist string
    	comma separated list of tables you want to generate dot file for
  -dbname string
    	dbname for which you want to generate dot file
  -sslmode
    	enable sslmode for postgres db connection
  -user string
    	username of postgres db
```

# TODO
Here's some features I would like to have for this project:
* Support connection string. What happens if the user doesn't have access to db?
* Add support for more db types.
* Prettify output.
* Add ability to whitelist columns in a table. Users should be able whitelist columns per table. But this starts getting into territory of language design. i.e. what kind of cli syntax should dbdot support. Hence this is the last item in my list.

# Inspiration
A while back I wanted a _simple_ tool that would just spit out schema for tables that I wanted. A lot of tools I found were way too powerful, requiring a zillion installation and configuration. This inspired me to write a simple self contained tool that was laser focused on just reading schema and spitting out DOT.
