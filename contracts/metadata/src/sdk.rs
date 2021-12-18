use serde::{Deserialize, Serialize};

pub fn focused_transform_patch_str(cid: &str, path: &str, prev: &str, next: &str) -> String {
    unsafe {
        let l = cid.len() as usize;
        let metadata = focused_transform_patch(cid, path, prev, next, NodeType::String).to_vec();
        String::from_utf8(metadata[..l].to_vec()).unwrap()
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

pub fn write_dag(data: &str) -> Vec<u8> {
    unsafe {
        let l = data.len() as usize;
        let metadata = write_dag_block(&data).to_vec();
        metadata[..l].to_vec()
    }
}
#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum NodeType {
    String = 0,
    Bytes = 1,
}

extern "C" {

    #[no_mangle]
    pub fn focused_transform_patch(
        cid: &str,
        path: &str,
        prev: &str,
        next: &str,
        ntype: NodeType,
    ) -> [u8; 1024];

    #[no_mangle]
    pub fn get_proof_by_cid(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn verify_proof_onchain(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn write_store(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn read_store(key: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn write_dag_block(data: &str) -> [u8; 1024];

    #[no_mangle]
    pub fn read_dag_block(cid: &str, ret: &i32) -> [u8; 1024];
}
