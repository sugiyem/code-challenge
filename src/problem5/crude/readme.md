# Crude
**Crude** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

## Functionalities

In this blockchain, user can do CRUD operation with the `Resource` object. Each `Resource` object will contains the following data:
- `id`, the unique identifier given to a resource when created
- `creator`, which refer to the resource's owner
- `metadata`, the information of resource (must be a string)
- `value`, the quantity of resource (must be a uint64)

Note that the `id` and `creator` can't be modified.

### Create Resource
The `create-resource` command will accept two inputs, string `metadata` and uint64 `value`.

For example, if we want to create a resource for Alice with `metadata=coin123` and `value=100`, simply run the command below.
```
cruded tx crude create-resource coin123 100 --from alice --chain-id crude
```
If we want to create another resource for Bob with `metadata=bobby dollar` and `value=1000`, we can execute the following command.
```
cruded tx crude create-resource "bobby dollar" 1000 --from bob --chain-id crude
```

### Show Resource Details
To view the details of a resource, we can use the `show-resource` command, which accept uint64 `id` as the input.

For instance if we want to view the resource with ID 1, we can simply run the following command.
```
cruded q crude show-resource 1
```
The result below will be shown in the CLI.
```
resource:
  creator: cosmos1mav60d2pnk6w5xlz3a60xyzwfnnfe90k3tcvnh
  id: "1"
  metadata: bobby dollar
  value: "1000"
```

### List Resources
The `list-resource` command will allow user to view filtered resources. This command expect three inputs, string `metadataFilter`, uint64 `valueLow`, and uint64 `valueHigh`. It will then print out all resources where:
- `metadataFilter` is the substring of `metadata` ; and
- `valueLow` <= `value` <= `valueHigh`

For example, if we want to list out all resources where the `value` is between 100 and 1000, inclusive, we can run the command
```
cruded q crude list-resource "" 100 1000
```
```
pagination:
  total: "2"
resource:
- creator: cosmos1w4j7j7ecyd42zfnkeama4v3zelzmgnn2qunecw
  id: "0"
  metadata: coin123
  value: "100"
- creator: cosmos1mav60d2pnk6w5xlz3a60xyzwfnnfe90k3tcvnh
  id: "1"
  metadata: bobby dollar
  value: "1000"
```

If we want to further filter the above resources such that the term `coin` must appear in `metadata`, we can run
```
cruded q crude list-resource coin 100 1000
```
```
pagination:
  total: "1"
resource:
- creator: cosmos1w4j7j7ecyd42zfnkeama4v3zelzmgnn2qunecw
  id: "0"
  metadata: coin123
  value: "100"
```


### Update Resource
The `update-resource` command will accept three inputs, string `new_metadata`, uint64 `new_value`, and uint64 `resource-id`. 
For instance, if we want to update the resource with ID 1 (which belongs to Bob) such that `new_metadata=bob yen` and `new_value=200`, we can run the following command.
```
cruded tx crude update-resource "bob yen" 200 1 --from bob --chain-id crude
```

Note that if we run the command below instead, the resource will not be updated as resource with ID 1 doesn't belong to Alice.
```
cruded tx crude update-resource "bob yen" 200 1 --from alice --chain-id crude
```

### Delete a Resource
To delete a specific resource, we can use the `delete-resource` command which accept uint64 `id` as the input.

For example, if we want to delete the resource with ID 0 (which belongs to Alice), we can run
```
cruded tx crude delete-resource 0 --from alice --chain-id crude
```
Similar to `update-resource`, the `delete-resource` can only be executed by the owner of the specific resource.
