use crate::verify_proof_onchain;
use crate::get_proof_by_cid;
use crate::sdk::focused_transform_patch_str;
use crate::sdk::read_dag;
use crate::sdk::{read_dag_block, write_dag_block};

extern crate juniper;

use juniper::{
    graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, GraphQLValue,
    RootNode, Variables,
};
use serde::{Deserialize, Serialize};

use std::collections::HashMap;

use std::str;
use std::vec::*;

pub struct Context {
    pub metadata: HashMap<String, Ancon721Metadata>,
}

impl juniper::Context for Context {}

#[derive(GraphQLObject, Clone, Debug, Serialize, Deserialize)]
pub struct DagLink {
    path: String,
    cid: String,
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct Ancon721Metadata {
    pub name: String,
    pub description: String,
    pub image: String,
    pub parent: String,
    pub owner: String,
    pub sources: Vec<String>,
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
pub struct Query;

#[graphql_object(context = Context)]
impl Query {
    fn api_version() -> &'static str {
        "0.1"
    }

    pub fn metadata(context: &Context, cid: String, path: String) -> Ancon721Metadata {
        let v = read_dag(&cid);
        let res = serde_json::from_slice(&v);
        res.unwrap()
    }
}

#[derive(Clone, Copy, Debug)]
pub struct Mutation;

#[graphql_object(context = Context)]
impl Mutation {
    async fn metadata(context: &Context, input: MetadataTransactionInput) -> Ancon721Metadata {
        let v = read_dag(&input.cid);
        let res = serde_json::from_slice(&v);
        let metadata: Ancon721Metadata = res.unwrap();
        let proof = get_proof_by_cid(&input.cid);
        let result = verify_proof_onchain(proof);
        if result.is_true() {
            let updated_cid =
                focused_transform_patch_str(&input.cid, "owner", &metadata.owner, &input.new_owner);
            let updated =
                focused_transform_patch_str(&updated_cid, "parent", &metadata.parent, &input.cid);

            let v = read_dag(&updated);
            let res = serde_json::from_slice(&v);
            let metadata = res.unwrap();
            apply_request_with_proof(tx, proof, offchain_data_cid, cid);
        }

        metadata
    }
}

#[derive(Clone, Debug, GraphQLInputObject)]
struct MetadataTransactionInput {
    path: String,
    cid: String,
    owner: String,
    new_owner: String,
}

type Schema = RootNode<'static, Query, Mutation, EmptySubscription<Context>>;

pub fn schema() -> Schema {
    Schema::new(Query, Mutation, EmptySubscription::<Context>::new())
}

pub fn apply_request_with_proof(tx: &str, proof: &str, cid: &str) -> &'static str {
    //  TODO:  submit_proof  and get_proof_by_cid
    ""
}
