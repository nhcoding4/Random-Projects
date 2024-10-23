# This is a reference implementation made by following a guide at
# https://medium.com/@vanflymen/learn-blockchains-by-building-one-117428612f46

# ----------------------------------------------------------------------------------------------------------------------

import json
from typing import List
from uuid import uuid4

from flask import Flask, jsonify, request

from blockchain import Blockchain


# ----------------------------------------------------------------------------------------------------------------------


def main() -> None:
    app = Flask(__name__)

    unique_identifier: str = str(uuid4()).replace("-", "")

    blockchain = Blockchain()

    # -----------------------------------------------------
    @app.route("/chain", methods=["GET"])
    def full_chain() -> tuple[json, int]:
        """ Displays the entire blockchain with each mined block and its transaction. """

        response: dict = {
            "chain": blockchain.chain,
            "length": len(blockchain.chain)
        }

        return jsonify(response), 200

    # -----------------------------------------------------

    @app.route("/transactions/new", methods=["POST"])
    def new_trans() -> tuple[str, int] | tuple[json, int]:
        """ Creates a new transaction to be mined into the Blockchain. Ensures the transaction is valid. """

        values: json = request.get_json()

        required_parameters: List[str] = ["sender", "recipient", "amount"]
        if not all(k in values for k in required_parameters):
            return "Missing values", 400

        index: int = blockchain.new_transaction(values["sender"], values["recipient"], values["amount"])
        response: json = {"message": f"Transaction will be added to Block {index}"}

        return jsonify(response), 201

    # -----------------------------------------------------

    @app.route("/mine", methods=["GET"])
    def mine() -> tuple[json, int]:
        """ Mines a block into the Blockchain. """

        # Run the proof of work algorithm.
        last_block: dict = blockchain.last_block
        last_proof: int = last_block["proof"]
        proof: int = blockchain.proof_of_work(last_proof)

        # Reward miner.
        blockchain.new_transaction(
            sender="0",
            recipient=unique_identifier,
            amount=1,
        )

        # Add block.
        previous_hash: str = blockchain.hash(last_block)
        block: dict = blockchain.new_block(proof, previous_hash)

        response = {
            "message": "New Block forged",
            "index": block["index"],
            "transactions": block["transactions"],
            "proof": block["proof"],
            "previous_hash": block["previous_hash"],
        }

        return jsonify(response), 200

    # -----------------------------------------------------

    @app.route("/nodes/register", methods=["POST"])
    def register_nodes() -> tuple[json, int] | tuple[str, int]:
        """ Registers a node into the network. """

        values: json = request.get_json()

        # No valid nodes.
        nodes: set[str] = values.get("nodes")
        if nodes is None:
            return "Error. Need a list of valid nodes", 400

        for node in nodes:
            blockchain.register_node(node)

        response: dict = {
            "message": "New nodes have been added",
            "total_nodes": list(blockchain.nodes),
        }
        return jsonify(response), 201

    # -----------------------------------------------------

    def consensus() -> tuple[json, int]:
        """ Allows the nodes to come to consensus on correct the blockchain. """

        replaced: bool = blockchain.resolve_conflicts()

        if replaced:
            response: dict = {
                "message": "Our chain was replaced",
                "new_chain": blockchain.chain
            }
        else:
            response = {
                "message": "Our chain is authoritative",
                "chain": blockchain.chain
            }

        return jsonify(response), 200

    # -----------------------------------------------------

    if __name__ == "__main__":
        app.run(host="0.0.0.0", port=5000)


# ----------------------------------------------------------------------------------------------------------------------


main()
