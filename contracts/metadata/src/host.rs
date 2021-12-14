use serde_json::json;
use std::collections::HashMap;
use crate::Context;
use juniper::Variables;
use crate::schema;
use wasm_bindgen::convert::IntoWasmAbi;
use wasm_bindgen::prelude::*;



#[wasm_bindgen()]
pub fn execute(query: String) -> String {
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
