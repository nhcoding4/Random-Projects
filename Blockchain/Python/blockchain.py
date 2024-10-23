# This is a reference implementation made by following a guide at
# https://medium.com/@vanflymen/learn-blockchains-by-building-one-117428612f46

# ---------------------------------------------------------------------------------------------------------------------

import hashlib
import json
from datetime import datetime
from typing import List
from urllib.parse import urlparse, ParseResult

import requests
from flask import Response


# ----------------------------------------------------------------------------------------------------------------------


class Blockchain(object):
    def __init__(self) -> None:
        self.chain: List[dict] = []
        self.current_transaction: List[dict] = []
        self.nodes: set[str] = set()

        # Genesis block:
        self.new_block(previous_hash="1", proof=100)

    # -----------------------------------------------------

    @property
    def last_block(self) -> dict:
        return self.chain[-1]

    # -----------------------------------------------------

    @staticmethod
    def hash(block: dict) -> str:
        """ Creates an SHA-256 hash of a Block """

        # Ensure the dictionary is sorted to prevent inconsistent hashes.
        block_string = json.dumps(block, sort_keys=True).encode()

        return hashlib.sha256(block_string).hexdigest()

    # -----------------------------------------------------

    @staticmethod
    def valid_proof(last_proof: int, proof: int) -> bool:
        """ Validates the proof. Does the passed parameters contain 4 leading 0's? """

        guess: bytes = f"{last_proof}{proof}".encode()
        guess_hash: str = hashlib.sha256(guess).hexdigest()
        return guess_hash[:4] == "0000"

    # -----------------------------------------------------

    def new_block(self, proof: int, previous_hash: str = None) -> dict:
        """ Creates a new block of transactions in the blockchain. """

        block: dict = {
            "index": len(self.chain) + 1,
            "timestamp": str(datetime.now()),
            "transactions": self.current_transaction,
            "proof": proof,
            "previous_hash": previous_hash or self.hash(self.chain[-1]),
        }

        self.current_transaction = []
        self.chain.append(block)
        return block

    # -----------------------------------------------------

    def new_transaction(self, sender: str, recipient: str, amount: int) -> int:
        """ Creates a new transaction to be mined into the next block """

        self.current_transaction.append({
            "sender": sender,
            "recipient": recipient,
            "amount": amount,
        })

        return self.last_block["index"] + 1

    # -----------------------------------------------------

    def proof_of_work(self, last_proof: int) -> int:
        """ Simple proof of work algorithm. """

        proof: int = 0

        while self.valid_proof(last_proof, proof) is False:
            proof += 1

        return proof

    # -----------------------------------------------------

    def register_node(self, address: str) -> None:
        """ Adds a new node to the list of nodes. """

        parsed_url: ParseResult = urlparse(address)
        self.nodes.add(parsed_url.netloc)

    # -----------------------------------------------------

    def valid_chain(self, chain: List[dict]) -> bool:
        """ Determine if a passed blockchain is valid """

        last_block: dict = chain[0]
        current_index: int = 1

        while current_index < len(chain):
            block: dict = chain[current_index]
            print(f"{last_block}")
            print(f"{block}")
            print("\n------------\n")

            if block["previous_hash"] != self.hash(last_block):
                return False

            if not self.valid_proof(last_block["proof"], block["proof"]):
                return False

            last_block = block
            current_index += 1

        return True

    # -----------------------------------------------------

    def resolve_conflicts(self) -> bool:
        """ Consensus Algorithm. Implements the longest chain as the canonical chain if there is a conflict. """

        neighbours: set[str] = self.nodes
        new_chain = None
        max_length: int = len(self.chain)

        # Verify the chains from all nodes in the network.
        for node in neighbours:
            response: Response = requests.get(f"http://{node}/chain")

            if response.status_code == 200:
                length: int = response.json()["length"]
                chain: List[dict] = response.json()["chain"]

                if length > max_length and self.valid_chain(chain):
                    max_length = length
                    new_chain = chain

        if new_chain:
            self.chain = new_chain
            return True

        return False
