use sdk::*;
mod sdk;
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
use serde::{Deserialize, Serialize};
use serde_json::json;

use std::collections::HashMap;
use std::env;

use std::str;
use std::vec::*;
use wasm_bindgen::prelude::*;

struct Context {
    metadata: HashMap<String, Ancon721Metadata>,
}

impl juniper::Context for Context {}

#[derive(GraphQLObject, Clone, Debug, Serialize, Deserialize)]
struct DagLink {
    path: String,
    cid: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
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
#[derive(Clone, Debug, Serialize, Deserialize)]
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
    fn api_version() -> &'static str {
        "0.1"
    }

    fn metadata(context: &Context, cid: String, path: String) -> Ancon721Metadata {
        unsafe {
            let m = read_dag_block(&cid, &path);
            let metadata = m.iter().map(|b| *b as char).collect::<String>();

            let cleaned = metadata.trim_end_matches(char::from(0));
            // let s = format!("{}", "{\"description\":\"description\",\"image\":\"http://ipfs.io/ipfs/\",\"name\":\"test\",\"owner\":\"\",\"parent\":\"\",\"sources\":[]}");
            let res = serde_json::from_str(&cleaned);
            res.unwrap()
        }
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

// #[derive(Clone, 000000Debug)]
// struct MetadataTransactionInput {
//   path: String,
//   cid: String,
//   owner: String,
//   newOwner: String,
// }

// #[derive(Clone, Debug)]
// struct Transaction {
//   metadata(tx: MetadataTransactionInput)-> JDagLink{}
// }
type Schema = RootNode<'static, Query, Mutation, EmptySubscription<Context>>;

fn schema() -> Schema {
    Schema::new(Query, Mutation, EmptySubscription::<Context>::new())
}

#[wasm_bindgen()]
pub fn execute(query: &str) -> String {
    // Create a context object.
    let ctx = Context {
        metadata: HashMap::default(),
    };

    let v = Variables::new();

    let sch = schema();

    let res = juniper::execute_sync(
        query, // "query { favoriteEpisode }",
        None, &sch, &v, &ctx,
    );
    let (data, err) = res.unwrap();
    let errors = err
        .iter()
        .map(|i| i.error().message().to_string())
        .collect::<Vec<String>>();

    json!({
        "data":data.to_string(),
        "errors": errors,
    })
    .to_string()
}

#[wasm_bindgen]
pub fn store(data: &str) -> Vec<u8> {
    let payload = Ancon721Metadata {
        name: "test".to_string(),
        description: "description".to_string(),
        image: "http://ipfs.io/ipfs/".to_string(),
        owner: "".to_string(),
        parent: "".to_string(),
        sources: [].to_vec(),
    };

    let json_payload = serde_json::to_string(&payload).unwrap();

    unsafe { write_dag_block(&json_payload).to_vec() }
}
