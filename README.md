# Ancon Hybrid Smart Contracts


## Install (Linux)

1. Install [WasmEdge](https://wasmedge.org/book/en/index.html).

2. Copy bin, include and lib to /usr/local

3. Configure rust with rustwasmc compiler

```
rustup default 1.50.0
rustup target add wasm32-wasi
curl https://raw.githubusercontent.com/second-state/rustwasmc/master/installer/init.sh -sSf | sh
rustwamc build
```

4. Now open repository in Visual Studio Code and run `main.go`.

## How to create hybrid smart contracts



The  hybrid smart contracts includes two directories: `adapters` where all secure offchain related  code is implemented, in Go.

>Note: Adapters must be compiled as a dependency together with Ancon Protocol Node.

The  `contracts` directory, each Rust Wasm library is added. Use the existing directory as boilerplate.


## API

### Graphql Code First 

Code example

```rust

// Enclosed with allow(unused) to enable WasmEdge and wasm_bindgen bindings
#![allow(unused)]
fn main() {
    use base64::*;
    use byteorder::{BigEndian, ReadBytesExt, WriteBytesExt};
    use hex::{FromHex, ToHex};

    // We are using Juniper as Graphql API
    use juniper::{
        graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, RootNode,
        Variables,
    };
    use serde_hex::utils::fromhex;
    use std::convert::TryInto;
    use std::fmt::Display;
    use std::io::Cursor;
    use std::str;
    use wasm_bindgen::prelude::*;

    #[derive(Clone, Copy, Debug)]
    struct Context;
    impl juniper::Context for Context {}

    #[derive(Clone, Copy, Debug, GraphQLEnum)]
    enum UserKind {
        Admin,
        User,
        Guest,
    }

    #[derive(Clone, Debug)]
    struct User {
        id: i32,
        kind: UserKind,
        name: String,
    }

    #[graphql_object(context = Context)]
    impl User {
        fn id(&self) -> i32 {
            self.id
        }

        fn kind(&self) -> UserKind {
            self.kind
        }

        fn name(&self) -> &str {
            &self.name
        }

        async fn friends(&self) -> Vec<User> {
            vec![]
        }
    }

    #[derive(Clone, Copy, Debug)]
    struct Query;

    #[graphql_object(context = Context)]
    impl Query {
        async fn users() -> Vec<User> {
            vec![User {
                id: 1,
                kind: UserKind::Admin,
                name: "user1".into(),
            }]
        }

        // WIP: Fetch from host function or instance
        async fn request(cid: String) -> Result<String, FieldError> {
            Ok(ancon::dag::get(&cid).await?.text().await?)
        }
    }

    type Schema = RootNode<'static, Query, EmptyMutation<Context>, EmptySubscription<Context>>;

    fn schema() -> Schema {
        Schema::new(
            Query,
            EmptyMutation::<Context>::new(),
            EmptySubscription::<Context>::new(),
        )
    }


    #[wasm_bindgen]
    pub fn execute(query: &str) -> JsFuture<Vec<String>> {
        // Create a context object.
        let ctx = Context;

        let m = EmptyMutation::new();
        let s = EmptySubscription::new();

        let v = Variables::new();

        let sch = Schema::new(Query, m, s);
        // Run the executor.
        let res = juniper::execute(
            query, // "query { favoriteEpisode }",
            None, &sch, &v, &ctx,
        );
        res
        // Ensure the value matches.
    }
}

```