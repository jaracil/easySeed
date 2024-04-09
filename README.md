# easySeed

EasySeed is a simple tool to generate a seed phrase for a wallet. It uses an easy to remember passPhrase to generate a seed phrase. The seed phrase is generated using the BIP39 standard.

It can generate 12 or 24 word seed phrases. The passPhrase can be any string of characters. The passPhrase is hashed using SHA256 to generate a 256 bit seed. The seed is then used to generate the seed phrase.

Don't use simple passPhrases. Use a passPhrase that is hard to guess. The passPhrase is the only thing that protects your seed phrase. If someone knows your passPhrase, they can generate your seed phrase and steal your funds.

