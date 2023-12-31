# SecureSeed
This program generates a secure seed from a 100 d6 dice throws - either provided by the user, or taken from [Random.Org](https://random.org).  
For a key to be secure, at least 100 bytes of entropy (well technically 99) are needed. These are provided as a string of digits 1-6, generate by dice throws. 

## Usage
Clone the repo, and either build, or run `app.go`.

- Use `-e` followed by your dice throws to provide your own entropy, or leave empty to get 100 random throws from Random.org.
- Use `-eth 3` to derive 3 Ethereum addresses from the seed
- Use `-btc 3` to derive 3 Bitcoin addresses from the seed
- use `-legacy` to derive a legacy (non-SegWit) Bitcoin addresses

```terminal
> go run app.go -eth 3 -e 3613462662326212532222252155632562142562242341312113253115434463256622124411233215321535122623166226
Dice throws:
3613462662326212532222252155632562142562242341312113253115434463256622124411233215321535122623166226

Private key:
06c603c87075b347634d4ba0c25b1343d78f59de590e2c3d3313c7ef9e5e7871

Mnemonic:
 1: almost       2: copy         3: velvet       4: thought 
 5: forest       6: photo        7: minor        8: practice
 9: patch       10: barrel      11: shadow      12: marble  
13: jump        14: provide     15: just        16: canyon  
17: club        18: trumpet     19: shaft       20: more    
21: wheat       22: connect     23: tide        24: derive  

Ethereum Addresses:
 1: 0x964C04Cd54f101d984B15fBf6E89553CFf401C10
 2: 0x3826b09ad6318C44bdF862A11416E1Eb634Bd165
 3: 0x908E9Be6EcB050dcbCD4da371E4daf2790dE66cc
```

## Using Random.org API
To use Random.org, please register on the site (free), create a free API key (1000 free calls supported), and place the URI and API key in a `.env` file at the top directory.
```env
RO_URI=https://api.random.org/json-rpc/4/invoke
RO_APIKEY=<your API key>
```

The URI above is the currently supported one (v4), and works well. Keep an eye out for changes.

If you do not feel like signing for the API, you can just use the site itself:
1. Free dice throws are limited to 60, so all you have to do is throw 50 twice
1. Go to [Random.org Dice Roller](https://www.random.org/dice/?num=50)
1. Write down the numbers - in any order you like
1. Refresh the page, and repeat
1. You now have 100 random digits to provide the program, for complete random entropy

## DIY
"Don't trust - verify" - yes, Random.org uses all kinds of nifty things like radio antennae listening to the radio waves in the atmosphere to generate randomness - but trusting your eyes is always better.

1. Get yourself a set of 10 d6 dice ($2 on Amazon, less on feeBay)
1. To ensure no "throw bias", put the dice is a transparent food container, and close it
1. Shake it off!
1. Write down the 10 numbers that came up
1. Go back to 3 (repeat 10 times)
1. You now have 100 random digits to provide the program, for complete random entropy
