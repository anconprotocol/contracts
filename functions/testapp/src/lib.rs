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
    pub extern "C" fn hello(input: &str) -> String {
        let output = "hola mundo";
        let res = to_hex(output.to_owned());
        output.to_owned()
    }

    fn to_hex(s: String) -> i64 {
        let bz = hex::encode(s);

        let mut bz = Cursor::new(bz);
        let res = bz.read_i64::<BigEndian>().unwrap_or_default();
        res
    }

    fn from_b64(v: i64) -> String {
        let mut args = vec![];
        args.write_i64::<BigEndian>(v.try_into().unwrap()).unwrap();
        let res = hex::decode(args).unwrap_or_default();
        String::from_utf8(res).unwrap()
    }

    // #[wasm_bindgen]
    // pub fn execute(query: &str) -> JsFuture<Vec<String>> {
    //     // Create a context object.
    //     let ctx = Context;

    //     let m = EmptyMutation::new();
    //     let s = EmptySubscription::new();

    //     let v = Variables::new();

    //     let sch = Schema::new(Query, m, s);
    //     // Run the executor.
    //     let res = juniper::execute(
    //         query, // "query { favoriteEpisode }",
    //         None, &sch, &v, &ctx,
    //     );
    //     res
    //     // Ensure the value matches.
    // }
}
