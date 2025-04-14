# automiesegnali

## Overview

The project `automiesegnali` studies the movement of point-like automata on uneven terrain, which are sensitive to signals attracting them to specific locations. It operates on a discrete grid plane defined by integer coordinates, employing a Manhattan distance metric for pathfinding and movement rules.

Key features include:
- Automata uniquely identified by binary strings.
- Obstacles defined as rectangular regions that limit movement.
- Signals emitted by sources to attract automata based on prefix matching.
- Movement algorithms ensuring only automata with minimum distance and free paths respond to signals.

This project integrates mathematical modeling, algorithmic design, and discrete geometry to explore automata behavior in constrained environments.

---

## Literate Programming

This project employs **Literate Programming**, an innovative programming paradigm introduced by Donald E. Knuth in his seminal work *Literate Programming* (1984). Literate Programming intertwines code and documentation, enabling developers to explain their logic and thought processes alongside the source code in a human-readable format.

The core of this project resides in the file `main.org`, which serves as the **original source**. Using tools inspired by the Literate Programming methodology, the following processes are automated:
1. **Tangling**: Extracting and generating the Go source code (`main.go`) from `main.org`.
2. **Weaving**: Generating an HTML file to provide a well-formatted, readable documentation.

This approach enhances clarity, maintainability, and collaboration by tightly coupling the implementation and its explanation.

## Repository Structure

- **`main.org`**: The primary Literate Programming source file. It contains both the documentation and the embedded Go code.
- **`main.go`**: Tangled Go code generated from `main.org`. This file contains the executable implementation.
- **`main.html`**: Weaved HTML documentation generated from `main.org`, providing a user-friendly visualization of the logic and code.

## Key Concepts

1. **Automata**:
   - Identified by binary strings (e.g., `1011`).
   - Positioned on a discrete grid with integer coordinates.
   - Capable of sharing positions with other automata.

2. **Obstacles**:
   - Defined as rectangular regions with integer boundaries.
   - Automata cannot occupy or traverse points within obstacles.

3. **Signals**:
   - Emit binary string signals to attract automata.
   - Automata respond if the signal is a prefix of their identifier.

4. **Movement Algorithm**:
   - Automata with minimum distance and free paths move towards the signal source.
   - Others remain stationary.

---

This project combines algorithmic rigor, discrete mathematics, and the elegance of Literate Programming to investigate automata behavior in complex environments. For more insights, explore the full documentation available in `main.html`.
