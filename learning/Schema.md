# Schema

#### 1. What is the smallest value a GraphQL `Int` may have?

- [ ] -9,223,372,036,854,775,808
- [ ] 2,147,483,648
- [ ] -2,147,483,647
- [ ] 0

<details closed>
<summary>🤨 Hint</summary>

A GraphQL `Int` is a signed 32 bit integer.

</details>

<details closed>
<summary>Answer</summary>
-2,147,483,648. The largest is 2,147,483,647.
</details>

#### 2. Which GraphQL type is used for decimals like `3.14`?

- [ ] `Decimal`
- [ ] `Float`
- [ ] `Number`
- [ ] `Double`

<details closed>
<summary>🤨 Hint</summary>
This type's name is often enjoyed with root beer and vanilla ice cream.
</details>

<details closed>
<summary>Answer</summary>

`Float`, capital F

</details>

#### 3. Which GraphQL type would be used for the value `"Foobar"`

- [ ] `[Char]`
- [ ] `[]Char`
- [ ] `string`
- [ ] `String`

<details closed>
<summary>🤨 Hint</summary>
This type is also a form of Cheese.
</details>

<details closed>
<summary>Answer</summary>

`String`, capital S

</details>

#### 4. Which `Boolean` value matches the statement, "I'm a good software engineer, capable of mastering GraphQL"?

<details closed>
<summary>🤨 Hint</summary>
believe in yourself! 🫶
</details>

<details closed>
<summary>Answer</summary>

`true`, lowercase t.

</details>

#### 5. While both are serialized to `String`s, what's the difference between the `ID` and `String` types?

<details closed>
<summary>🤨 Hint</summary>

Example values of `ID`s may include uuids or hashes, like `623584F7-C060-4142-8223-5E8A6D066CF2` or `d732fee6462de7f04f9432f1bb3925f57554db1d8c8d6f3138eea70e5787c7ae`.

Example values of `String`s may include: `"I'm human-readable!"` or `"GraphQL IDs aren't intended to be human-readable!"`

</details>

<details closed>
<summary>Answer</summary>

GraphQL `ID`s aren't intended to be human-readable. They are often used as unique identifiers for client-side caching.

</details>

#### 6. `enum`s are a special object type that signify a finite set states.

- [ ] True
- [ ] False

<details closed>
<summary>🤨 Hint</summary>

Objects have multiple fields with different values. Scalars can hold one value. A declared `enum` may offer multiple values, but how many of those values can a `enum` take in a query?

</details>

<details closed>
<summary>Answer</summary>

False. `enum`s are a special **scalar** type that signify a finite set of states.

</details>

### 7. Which fields of `Node` are nullable?

```graphql
# Node is the interface for something in the file system.
interface Node {
  id: ID!
  name: String!
  owner: User!
  parent: Folder
}
```

<details closed>
<summary>🤨 Hint</summary>

Unless followed by a `!`, all types are nullable.

</details>

<details closed>
<summary>Answer</summary>

`Parent` is the only nullable field as it's the only field of `Node` that is not modified by a suffixed `!`.

</details>

### 8. What is the type of `Folder.children`?

```graphql
# Folder is a node of the file system that holds other nodes.
type Folder implements Node {
  # Node fields
  id: ID!
  name: String!
  owner: User!
  parent: Folder

  # Folder field
  children: [Node!]!
}
```

<details closed>
<summary>🤨 Hint</summary>

`[Int]` signifies a list of integers.

</details>

<details closed>
<summary>Answer</summary>

`[Node!]!` signifies a non-nullable list of non-nullable `Node`s.

</details>

### 9. `interface`s are useful when...

- [ ] You want to return an object or set of objects, but those might be of several different types.
- [ ] You want to define the behavior of a object, but not its values.
- [ ] You want users to write a bunch of separate queries that are very similar, but only differ on the object(s) that are returned.
- [ ] All of the above.

<details closed>
<summary>🤨 Hint</summary>

It's not all of the above.

</details>

<details closed>
<summary>Answer</summary>

The first option. `interface`s allow users to write flexible queries that request `interace`s and resolve to concrete types using inline fragments.

Consider the schema:

```graphql
# Node is the interface for something in the file system.
interface Node {
  id: ID!
  name: String!
  owner: User!
  parent: Folder
}

# Folder is a node of the file system that holds other nodes.
type Folder implements Node {
  # Node fields
  id: ID!
  name: String!
  owner: User!
  parent: Folder

  # Folder field
  children: [Node!]!
}

# File is a node of the file system that contents content.
type File implements Node {
  # Node fields
  id: ID!
  name: String!
  owner: User!
  parent: Folder

  # File field
  content: String!
}

# Query is the root query type.
#
# It defines the interface in which clients will query the shared file system with.
type Query {
  node(id: ID): Node
}
```

Now, see the query:

```graphql
query {
  node {
    ... on Folder {
      children {
        id
      }
    }
    ... on File {
      content
    }
  }
}
```

Using the `Node` interface, we give clients one query to use when they want to find either a `Folder` or `File`.

</details>

### 10. What are the two special object types that act as the entrypoint to a GraphQL service?

<details closed>
<summary>🤨 Hint</summary>

These two object types allow clients to _query_ and _mutate_ a GraphQL service.

</details>

<details closed>
<summary>Answer</summary>

The types `Query` and `Mutation` are the client's entrypoints to querying and mutating a GraphQL service, respectively.

</details>
