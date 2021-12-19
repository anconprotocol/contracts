use serde::{Deserialize, Serialize};
use serde_json::json;

pub fn focused_transform_patch_str(cid: &str, path: &str, prev: &str, next: &str) -> String {
    unsafe {
        let s: i32 = 0;
        let payload = json!({
            "cid": cid,
            "nodeType": NodeType::String,
            "previousValue": prev,
            "nextValue": next,
            "path": path
        });
        let lnk = focused_transform_patch(&payload.to_string(), &s).to_vec();
        let x = s as usize;
        let v = &lnk[..x];
        String::from_utf8(v.to_vec()).unwrap()
    }
}

pub fn read_dag(cid: &str) -> Vec<u8> {
    unsafe {
        let s: i32 = 0;
        let metadata = read_dag_block(&cid, &s);
        let x = s as usize;
        let v = &metadata[..x];
        v.to_vec()
    }
}

pub fn submit_proof(payload: &str, prev_proof: &str, new_cid: &str) -> String {
    unsafe {
        let l = payload.len() as usize;
        let res = submit_proof_onchain(&payload, prev_proof, new_cid).to_vec();
        String::from_utf8(res[..l].to_vec()).unwrap()
    }
}

pub fn get_proof(cid: &str) -> String {
    unsafe {
        let s: i32 = 0;
        let proof = get_proof_by_cid(&cid, &s);
        let x = s as usize;
        let v = &proof[..x];
        String::from_utf8(v.to_vec()).unwrap()
    }
}

pub fn verify_proof(data: &str) -> bool {
    unsafe {
        let s: i32 = 0;
        let res = verify_proof_onchain(&data, &s);
        let x = s as usize;
        let v = &res[..x];
        let ok = String::from_utf8(v.to_vec()).unwrap();
        ok == "true"
    }
}

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum NodeType {
    String = 0,
    Bytes = 1,
}

extern "C" {

    #[no_mangle]
    pub fn submit_proof_onchain(input: &str, prev_proof: &str, cid: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn focused_transform_patch(args: &str, ret: &i32) -> [u8; 1024];

    #[no_mangle]
    pub fn get_proof_by_cid(key: &str, ret: &i32) -> [u8; 1024];

    #[no_mangle]
    pub fn verify_proof_onchain(key: &str, ret: &i32) -> [u8; 1024];

    #[no_mangle]
    pub fn write_store(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn read_store(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn write_dag_block(data: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn read_dag_block(cid: &str, ret: &i32) -> [u8; 1024];
}
