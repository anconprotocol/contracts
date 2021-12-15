use ssvm_wasi_helper::ssvm_wasi_helper::_initialize;

extern crate juniper;

#[macro_use]
extern crate juniper_codegen;
use std::env;
use base64::*;
use byteorder::{BigEndian, ReadBytesExt, WriteBytesExt};
use hex::{FromHex, ToHex};
use juniper::{
    graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, GraphQLValue,
    RootNode, Variables,
};
use serde_json::json;
use std::collections::HashMap;

use serde_hex::utils::fromhex;
use std::convert::TryInto;
//    use std::convert::From::from;
use std::fmt::Display;
use std::future::*;
use std::io::Cursor;
use std::str;
use std::vec::*;
use wasm_bindgen::prelude::*;

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
    
    fn api_version() -> &'static str {
        "0.1"
    }

    fn metadata(context: &Context,cid: String, path: String) -> Ancon721Metadata {

        
unsafe {let metadata = read_dag_block(&cid, &path);}

       Ancon721Metadata {
            name: "test".to_string(),
            description: "description".to_string(),
            image: "http://ipfs.io/ipfs/".to_string(),
            owner: "".to_string(),
            parent: "".to_string(),
            sources: [].to_vec(),
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
    }).to_string()
}
extern    {
   // #[no_mangle]
    pub fn write_store(key: String, value: String);

   // #[no_mangle]
    pub fn read_store(key: String) -> String;

  //  #[no_mangle]
  pub fn write_dag_block(data: String) -> String;

    #[no_mangle]
    pub fn read_dag_block(cid: &str, path: &str) -> String; 
}
