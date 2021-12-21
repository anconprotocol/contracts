use crate::sdk::focused_transform_patch_str;
use crate::sdk::read_dag;
use crate::sdk::submit_proof;
use crate::sdk::{generate_proof, get_proof, read_dag_block, write_dag_block};
use juniper::FieldResult;

extern crate juniper;

use juniper::{
    graphql_object, EmptyMutation, EmptySubscription, FieldError, GraphQLEnum, GraphQLValue,
    RootNode, Variables,
};
use serde::{Deserialize, Serialize};
use serde_json::json;

use std::collections::HashMap;

use std::str;
use std::vec::*;

pub struct Context {
    pub metadata: HashMap<String, Ancon721Metadata>,
    pub transfer: HashMap<String, MetadataPacket>,
}

impl juniper::Context for Context {}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub struct MetadataPacket {
    pub cid: String,
    pub from_owner: String,
    pub result_cid: String,
    pub to_owner: String,
    pub to_address: String,
    pub token_id: String,
    pub proof: String,
}

#[graphql_object(context = Context)]
impl MetadataPacket {
    fn cid(&self) -> &str {
        &self.cid
    }

    fn from_owner(&self) -> &str {
        &self.from_owner
    }

    fn result_cid(&self) -> &str {
        &self.result_cid
    }
    fn to_owner(&self) -> &str {
        &self.to_owner
    }

    fn to_address(&self) -> &str {
        &self.to_address
    }

    fn token_id(&self) -> &str {
        &self.token_id
    }
    fn proof(&self) -> &str {
        &self.proof
    }
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
        &self.owner
    }

    async fn sources(&self) -> &Vec<String> {
        &self.sources
    }
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
    //Dagblock mutation
    fn transfer(context: &Context, input: MetadataTransactionInput) -> Vec<MetadataPacket> {
        let v = read_dag(&input.cid);
        let res = serde_json::from_slice(&v);
        let metadata: Ancon721Metadata = res.unwrap();

        //generate current metadata proof packet
        let proof = generate_proof(&input.cid);

        let updated_cid =
            focused_transform_patch_str(&input.cid, "owner", &metadata.owner, &input.new_owner);
        let updated =
            focused_transform_patch_str(&updated_cid, "parent", &metadata.parent, &input.cid);

        //generate updated metadata proof packet
        let proof_cid = apply_request_with_proof(input, &proof, &updated);
        let v = read_dag(&proof_cid);
        let res = serde_json::from_slice(&v);
        let packet: MetadataPacket = res.unwrap();
        let current_packet = MetadataPacket {
            cid: input.cid,
            from_owner: val,
            result_cid: updated,
            to_owner: val,
            to_address: val,
            token_id: val,
            proof: proof,
        };
        let result = vec![current_packet, packet];
        result
    }
}

#[derive(Clone, Debug, GraphQLInputObject, Serialize, Deserialize)]
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

fn apply_request_with_proof(
    input: MetadataTransactionInput,
    prev_proof: &str,
    new_cid: &str,
) -> String {
    // Must combined proofs (prev and new) in host function
    // then send to chain and return result
    let js = json!({
        "previous": prev_proof,
        "next_cid": new_cid,
        "inpur": input
    });
    submit_proof(&js.to_string())
}
