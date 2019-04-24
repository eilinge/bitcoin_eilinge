entry
    coin/main
call
    core
        block
            block struct
            newblock(string, []byte)
            setHash
            newGenesisBlock
        blockchain
            addBlock
            newBlockChain
        proofOfWork
            newProofOfWork(b)
            perpareData(nonce)
            Run: get nonce
            Validate
        utils
            IntToHex
            DataToHex
