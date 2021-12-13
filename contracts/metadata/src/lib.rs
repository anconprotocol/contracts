#![allow(unused)]
fn main() {
    use base64::*;
    use byteorder::{BigEndian, ReadBytesExt, WriteBytesExt};
    use hex::{FromHex, ToHex};
    use juniper::{
        graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, RootNode,
        Variables,
    };
    use serde_hex::utils::fromhex;
    use std::convert::TryInto;
//    use std::convert::From::from;
    use std::fmt::Display;
    use std::io::Cursor;
    use std::future::*;
    use std::str;
    use wasm_bindgen_futures::*;
    use wasm_bindgen::prelude::*;
    use wasm_bindgen::convert::IntoWasmAbi;
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

        // /// Fetch a URL and return the response body text.
        // async fn request(url: String) -> Result<String, FieldError> {
        //     Ok(reqwest::get(&url).await?.text().await?)
        // }
    }

    type Schema = RootNode<'static, Query, EmptyMutation<Context>, EmptySubscription<Context>>;

    fn schema() -> Schema {
        Schema::new(
            Query,
            EmptyMutation::<Context>::new(),
            EmptySubscription::<Context>::new(),
        )
    }
    fn main() {}

    #[wasm_bindgen]
    extern "C" {
        #[wasm_bindgen]
        pub fn response(s: &str) -> String;
        #[wasm_bindgen]
        pub fn args() -> String;
    }

    #[wasm_bindgen]
    pub fn echo(content: &str) -> String {
        println!("Printed from wasi: {}", content);
        return content.to_string();
    }

    #[wasm_bindgen]
    pub fn hello(input: &str) -> (String) {
        let output = "hola mundo";
        let res = hex::encode(output.to_owned());
        res
    }


    #[wasm_bindgen]
    pub async fn execute(query: String) -> js_sys::Promise {
        // Create a context object.
        let ctx = Context;

        let m = EmptyMutation::new();
        let s = EmptySubscription::new();

        let v = Variables::new();

        let sch = Schema::new(Query, m, s);
        // Run the executor.
        let res = juniper::execute(
            &query, // "query { favoriteEpisode }",
            None, &sch, &v, &ctx,
        ).await;
        let (data, err) = res.unwrap();

        let x=data.to_string();
        let promise = js_sys::Promise::resolve(&x.into());

        // let result = wasm_bindgen_futures::JsFuture::from(promise).await?;
        // Ok(result)
        promise
        // Ensure the value matches.
    }
}
