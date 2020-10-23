package hello;

public class Greeter {
  public String sayHello() {
    return "Hello world!";
  }
  public String getEnv() {
    String envv = System.getenv("TEST");
    return envv;
  }
}