# automiesegnali

The problem involves:

    Studying the movement of point-like automata on uneven terrain that are sensitive to signals attracting them to certain locations.
    Working on a plane defined as a set of points with integer coordinates: {(x, y) ∈ Z × Z}
    Defining movements in terms of:
        Horizontal unit steps: segments connecting points (x, y) and (x + 1, y)
        Vertical unit steps: segments connecting points (x, y) and (x, y + 1)
    A path is defined as a sequence S = p1, p2, ..., pk of k horizontal and/or vertical unit steps where:
        Consecutive steps share only one vertex
        No point is common to more than two unit steps
        The length of S is k
    The distance D(A, B) between two points A and B is defined as the minimum length of paths connecting them
    For points A(xA, yA) and B(xB, yB), the distance is calculated as:
    D(A, B) = |xB - xA| + |yB - yA|

This is essentially describing a Manhattan distance (or L1 distance) metric on a discrete grid.

Key points about the automata:

    Each automaton is uniquely identified by a name η (eta)
    The name η is a finite string over the alphabet {0, 1}
        η = b1b2...bn where n is a positive integer
        Each bi ∈ {0, 1} for i ∈ {1, ..., n}
    The position P(η) of an automaton at any given time is specified by coordinates (x0, y0) ∈ Z × Z
    Multiple automata can occupy the same point on the plane

Key points about obstacles:

    Obstacles are present in the plane and limit the movement of automata
    Each obstacle is defined as a set of points contained within a rectangle with vertical and horizontal sides
    A rectangle is denoted as R(x0, y0, x1, y1) where:
        (x0, y0) are coordinates of the bottom-left vertex
        (x1, y1) are coordinates of the top-right vertex
        All coordinates are integers
    Important rules:
        Obstacles can overlap
        No automaton can be positioned on any point belonging to an obstacle
        A path is considered "free" if it doesn't intersect with any obstacles

Key points about signals and movement:

    Sources can emit recall signals (α) which are finite strings on the alphabet {0, 1}
    An automaton (η) responds to a signal (α) if and only if α is a prefix of η
        For example, if η = "1011", it responds to signals "1", "10", "101", "1011"
    Movement rules:
        Among responding automata, only those with minimum distance from the source will move
        These automata must also have a free path of minimum distance available
        All qualifying automata will move to the source position

Let's break down the movement algorithm:

    When a source at position (x, y) emits signal α:
        Find all automata {η1, ..., ηk} where α is a prefix of ηi
        Calculate distance di from each automaton to source
        Find minimum distance d = min{di | 1 ≤ i ≤ k}
    Define set A as automata that:
        Have distance = d from source
        Have a free path of length d to source
    Movement:
        All automata in set A move to source position (x, y)
        All other automata remain in their current positions
