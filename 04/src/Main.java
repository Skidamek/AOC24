import java.io.FileNotFoundException;
import java.io.FileReader;
import java.util.ArrayList;
import java.util.List;

public class Main {

    public static List<List<Character>> lines = new ArrayList<>(); // y, x
    public static int gridMaxX;
    public static int gridMaxY;

    public static void main(String[] args) throws FileNotFoundException {

        long startTime = System.currentTimeMillis();

        FileReader fileReader = new FileReader("../input.txt");

        try {
            int i;
            List<Character> line = new ArrayList<>();
            while ((i = fileReader.read()) != -1) {
                char ch = (char) i;
                if (ch == '\n') {
                    lines.add(line);
                    line = new ArrayList<>();
                } else {
                    line.add(ch);
                }
            }
            lines.add(line);
        } catch (Exception e) {
            e.printStackTrace();
        }

        gridMaxX = lines.stream().mapToInt(List::size).max().orElse(0) - 1;
        gridMaxY = lines.size() - 1;

        List<Point[]> xmases1 = part1();
        List<Point[]> xmases2 = part2();

        long elapsedTime = System.currentTimeMillis() - startTime;
//        printGrid(xmases1);

        System.out.println("Part 1 - XMAS count: " + xmases1.size());

//        printGrid(xmases2);

        System.out.println("Part 2 - XMAS count: " + xmases2.size() / 2);

        System.out.println("Took: " + (elapsedTime * 1000) + "µs");
    }

    public static List<Point[]> part2() {
        List<Point[]> xmases = new ArrayList<>();

        // search XMAS
        for (int y = 0; y < lines.size(); y++) {
            List<Character> line = lines.get(y);
            for (int x = 0; x < line.size(); x++) {
                char ch = line.get(x);
                if (ch != 'A') {
                    continue;
                }

                Point X = null;
                Point M = null;
                Point A = new Point(x, y);
                Point S = null;

                List<Point> possibleMs = new ArrayList<>();
                List<Point> possibleSs = new ArrayList<>();

                // search for M and S around the A
                for (int wx = -1; wx < 2; wx++) {
                    for (int wy = -1; wy < 2; wy++) {
                        if (wx == 0 && wy == 0) {
                            continue;
                        }

                        int xx = x + wx;
                        int yy = y + wy;

                        if (xx < 0 || xx > gridMaxX || yy < 0 || yy > gridMaxY) {
                            continue;
                        }

                        char c = lines.get(yy).get(xx);
                        switch (c) {
                            case 'M':
                                possibleMs.add(new Point(xx, yy));
                                break;
                            case 'S':
                                possibleSs.add(new Point(xx, yy));
                                break;
                        }
                    }
                }

                // match M and S
                for (Point possibleM : possibleMs) {
                    for (Point possibleS : possibleSs) {
                        if (A.crossMiddlePointOf(possibleM, possibleS)) {
                            M = possibleM;
                            S = possibleS;


                            xmases.add(new Point[] {M, A, S});
                        }
                    }
                }
            }
        }

        // remove X-MAS which are not crossing any other X-MAS
        List<Point[]> xmasesToRemove = new ArrayList<>();
        for (Point[] xmas : xmases) {
            Point M = xmas[0];
            Point A = xmas[1];
            Point S = xmas[2];

            boolean isCrossing = false;
            for (Point[] xmas2 : xmases) {
                if (xmas == xmas2) {
                    continue;
                }

                Point M2 = xmas2[0];
                Point A2 = xmas2[1];
                Point S2 = xmas2[2];

                if (A.crossMiddlePointOf(M2, S2)) {
                    isCrossing = true;
                    break;
                }
            }

            if (!isCrossing) {
                xmasesToRemove.add(xmas);
            }
        }

        xmases.removeAll(xmasesToRemove);

        return xmases;
    }

    public static List<Point[]> part1() {
        List<Point[]> xmases = new ArrayList<>();

        // search XMAS
        for (int y = 0; y < lines.size(); y++) {
            List<Character> line = lines.get(y);
            for (int x = 0; x < line.size(); x++) {
                char ch = line.get(x);
                if (ch != 'A') {
                    continue;
                }

                Point X = null;
                Point M = null;
                Point A = new Point(x, y);
                Point S = null;

                List<Point> possibleMs = new ArrayList<>();
                List<Point> possibleSs = new ArrayList<>();

                // search for M and S around the A
                for (int wx = -1; wx < 2; wx++) {
                    for (int wy = -1; wy < 2; wy++) {
                        if (wx == 0 && wy == 0) {
                            continue;
                        }

                        int xx = x + wx;
                        int yy = y + wy;

                        if (xx < 0 || xx > gridMaxX || yy < 0 || yy > gridMaxY) {
                            continue;
                        }

                        char c = lines.get(yy).get(xx);
                        switch (c) {
                            case 'M':
                                possibleMs.add(new Point(xx, yy));
                                break;
                            case 'S':
                                possibleSs.add(new Point(xx, yy));
                                break;
                        }
                    }
                }

                // match M and S
                for (Point possibleM : possibleMs) {
                    for (Point possibleS : possibleSs) {
                        if (A.middlePointOf(possibleM, possibleS)) {
                            M = possibleM;
                            S = possibleS;
                            X = M.getOppositePoint(A);

                            if (X == null) {
                                continue;
                            }
                            Character xChar = X.getChar(lines);
                            if (xChar == null) {
                                continue;
                            }
                            if (xChar == 'X') {
                                xmases.add(new Point[] {X, M, A, S});
                            }
                        }
                    }
                }
            }
        }

        return xmases;
    }

    public static void printGrid(List<Point[]> xmases) {
        char[][] grid = new char[gridMaxY + 1][gridMaxX + 1];
        for (int y = 0; y <= gridMaxY; y++) {
            for (int x = 0; x <= gridMaxX; x++) {
                grid[y][x] = '.';
            }
        }

        xmases.forEach(xmas -> {
            for (Point point : xmas) {
                grid[point.y][point.x] = point.getChar(lines);
            }
        });

        for (char[] row : grid) {
            for (char cell : row) {
                System.out.print(cell);
            }
            System.out.println();
        }
    }

    public record Point(int x, int y) {
        public Point getOppositePoint(Point nextPoint) {
            int deltaX = nextPoint.x - this.x;
            int deltaY = nextPoint.y - this.y;

            // Opposite point is calculated by reversing the direction
            int oppositeX = this.x - deltaX;
            int oppositeY = this.y - deltaY;

            if (oppositeX < 0 || oppositeY < 0) {
                return null;
            }

            return new Point(oppositeX, oppositeY);
        }

        public boolean middlePointOf(Point pointOne, Point pointTwo) {
            int middleX = (pointOne.x + pointTwo.x) / 2;
            int middleY = (pointOne.y + pointTwo.y) / 2;
            if (middleX == this.x && middleY == this.y) {
                // check if pointOne x/y and pointTwo x/y is only one x/y aprt from the middle point so its in the valid line
                for (int x = -1; x < 2; x++) {
                    for (int y = -1; y < 2; y++) {
                        if (pointOne.x == middleX + x && pointOne.y == middleY + y) {
                            if (pointTwo.x == middleX - x && pointTwo.y == middleY - y) {
                                return true;
                            }
                        }
                    }
                }
            }
            return false;
        }

        public boolean crossMiddlePointOf(Point pointOne, Point pointTwo) {
            if (pointOne.x == pointTwo.x || pointOne.y == pointTwo.y) {
                return false;
            }

            return middlePointOf(pointOne, pointTwo);
        }

        public Character getChar(List<List<Character>> lines) {
            if (y < 0 || y >= lines.size() || x < 0 || x >= lines.get(y).size()) {
                return null;
            }
            return lines.get(y).get(x);
        }

        public String toString() {
            return String.format("(x = %d, y= %d)", x , y);
        }
    }
}