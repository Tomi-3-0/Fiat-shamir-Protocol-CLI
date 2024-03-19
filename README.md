## Fiat-Shamir heuristic Protocol

### Introduction
    A cryptographic protocol is a distributed algorithm describing precisely the interactions between two or more entities, achieving certain security objectives.

    There are basically two forms of authentication schemes: interactive authentication schemes (e.g., identification schemes) and noninteractive authentication schemes (e.g., digital signature schemes). Similarly, one may distinguish two forms of zero-knowledge proof schemes: interactive proof schemes and noninteractive proof schemes.

    An interactive proof scheme comprises a protocol by which a prover convinces a verifier that a certain statement holds. A noninteractive proof scheme comprises an algorithm by which a prover generates a proof for a certain statement and another algorithm by which a verifier may verify a given proof. It turns out that there is a simple but effective way to make any Σ-protocol noninteractive, known as the Fiat–Shamir heuristic. 
    A distinctive feature of a noninteractive Σ-proof is that any entity may play the role of the verifier. As a consequence, a noninteractive Σ-proof can be verified independently by many entities—just as a digital signature can be verified by anyone who is interested in its validity.

    A practical scheme for zero-knowledge identification was first presented by Fiat and Shamir [FS87], followed by the slightly improved scheme of Feige, Fiat, and Shamir

### References
Schoenmakers, B. (2024, February 2). Lecture Notes Cryptographic Protocols. https://www.win.tue.nl/~berry/CryptographicProtocols/LectureNotes.pdf.
U. Feige, A. Fiat, and A. Shamir. Zero-knowledge proofs of identity. Journal of Cryptology, 1(2):77–94, 1988.