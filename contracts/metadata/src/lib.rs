extern crate juniper;

#[macro_use]
extern crate juniper_codegen;

use base64::*;
use byteorder::{BigEndian, ReadBytesExt, WriteBytesExt};
use hex::{FromHex, ToHex};

use juniper::{
    graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, GraphQLValue,
    RootNode, Variables,
};
use std::collections::HashMap;

use serde_hex::utils::fromhex;
use std::convert::TryInto;
//    use std::convert::From::from;

use std::str;

struct Context {
    metadata: HashMap<String, Ancon721Metadata>,
}

impl juniper::Context for Context {}

#[derive(GraphQLObject, Clone, Debug)]
struct DagLink {
    path: String,
    cid: String,
}

#[derive(Clone, Debug)]
struct Ancon721Metadata {
    name: String,
    description: String,
    image: String,
    parent: String,
    owner: String,
    sources: Vec<String>,
}

#[graphql_object(context = Context)]
impl Ancon721Metadata {
    fn name(&self) -> &str {
        &self.name
    }

    fn description(&self) -> &str {
        &self.description
    }

    fn image(&self) -> &str {
        &self.image
    }
    fn parent(&self) -> &str {
        &self.parent
    }

    fn owner(&self) -> &str {
        &self.parent
    }

    async fn sources(&self) -> Vec<String> {
        vec![]
    }


}
#[derive(Clone, Debug)]
struct DagContractTrusted {
    data: DagLink,
    payload: Ancon721Metadata,
}

// pub struct Subscription;

// type StringStream = Pin<Box<dyn Stream<Item = Result<String, FieldError>> + Send>>;

// #[graphql_subscription(context = Database)]
// impl Subscription {
//     async fn hello_world() -> StringStream {
//         let stream =
//             futures::stream::iter(vec![Ok(String::from("Hello")), Ok(String::from("World!"))]);
//         Box::pin(stream)
//     }
// }
#[derive(Clone, Copy, Debug)]
struct Query;

#[graphql_object(context = Context)]
impl Query {
    async fn metadata(cid: String, path: String) -> Vec<Ancon721Metadata> {
      //  let metadata = host::read_dag_block(cid, path);

        vec![Ancon721Metadata {
            name: "test".to_string(),
            description: "description".to_string(),
            image: "http://ipfs.io/ipfs/".to_string(),
            owner: "".to_string(),
            parent: "".to_string(),
            sources: [].to_vec(),
        }]
    }
}

#[derive(Clone, Copy, Debug)]
struct Mutation;

#[graphql_object(context = Context)]
impl Mutation {
    async fn metadata(cid: String, path: String) -> Vec<Ancon721Metadata> {
        vec![Ancon721Metadata {
            name: "test".to_string(),
            description: "description".to_string(),
            image: "http://ipfs.io/ipfs/".to_string(),
            owner: "".to_string(),
            parent: "".to_string(),
            sources: [].to_vec(),
        }]
    }

    // /// Fetch a URL and return the response body text.
    // async fn request(url: String) -> Result<String, FieldError> {
    //     Ok(reqwest::get(&url).await?.text().await?)
    // }
}

// #[derive(Clone, Debug)]
// struct MetadataTransactionInput {
//   path: String,
//   cid: String,
//   owner: String,
//   newOwner: String,
// }

// #[derive(Clone, Debug)]
// struct Transaction {
//   metadata(tx: MetadataTransactionInput)-> DagLink{}
// }
type Schema = RootNode<'static, Query, Mutation, EmptySubscription<Context>>;

fn schema() -> Schema {
    Schema::new(Query, Mutation, EmptySubscription::<Context>::new())
}

pub async fn execute(query: String) -> js_sys::Promise {
    // Create a context object.
    let ctx = Context { metadata: HashMap::default() };

    let s = EmptySubscription::new();
    let v = Variables::new();

    let sch = Schema::new(Query, Mutation, s);
    // Run the executor.
    let res = juniper::execute(
        &query, // "query { favoriteEpisode }",
        None, &sch, &v, &ctx,
    )
    .await;
    let (data, err) = res.unwrap();

    let x = data.to_string();
    let promise = js_sys::Promise::resolve(&x.into());

    // let result = wasm_bindgen_futures::JsFuture::from(promise).await?;
    // Ok(result)
    promise
    // Ensure the value matches.
}

fn main() {}
