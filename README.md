# easySeed

EasySeed is a simple tool to generate a deterministic seed phrase for a wallet. It uses an easy to remember passPhrase to generate a seed phrase. The seed phrase is generated using the BIP39 standard.

It can generate 12 or 24 word seed phrases. The passPhrase can be any string of characters. The passPhrase is hashed using SHA256 to generate a 256 bit seed. The seed is then used to generate the seed phrase.

Don't use simple passPhrases. Use a passPhrase that is hard to guess. The passPhrase is the only thing that protects your seed phrase. If someone knows your passPhrase, they can generate your seed phrase and steal your funds.

You can download binary files here: https://github.com/jaracil/easySeed/releases/

You can create a memorable passphrase based on a personal phrase that only you would know, combining it with details that are not obvious. For example:

        Take a meaningful phrase: "My dog Bruno was born in 2010"
        Modify it to add complexity: "MyDogBrunoBrunoBornIn2010!"

        You can derivate other wallets from the same passPhrase by using the same passPhrase with a suffix. For example:

        PassPhrase: "MyDogBrunoBrunoBornIn2010!"
        Wallet 1: "MyDogBrunoBrunoBornIn2010!-wallet1"
        Wallet 2: "MyDogBrunoBrunoBornIn2010!-wallet2"
        ... 