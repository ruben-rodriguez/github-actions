package hello;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

public class HelloWorld {
  public static void main(String[] args) {
    Greeter greeter = new Greeter();
    System.out.println(greeter.sayHello());
    String env = greeter.getEnv();
    System.out.println(env);
  }
}