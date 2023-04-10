public class Main {
    public static void main(String[] args) {
        String name , age ;
        name = "Hello Wrold";
        age = "25";
        int i = 0;
        while (i < 5) {
          System.out.println(i + name + age);
          i++;
        }

        System.out.println();
        for (i = 0; i<5; i++) {
            System.out.println(i + name + age);
        }

        String[] cars = {"a", "b", "c" , "d"};

        for (String c : cars) {
            System.out.println(cars.length);
            System.out.println(c);
        }
     
    }
  }