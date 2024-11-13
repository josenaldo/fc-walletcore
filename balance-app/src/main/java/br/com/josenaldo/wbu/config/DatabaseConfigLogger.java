package br.com.josenaldo.wbu.config;
import jakarta.annotation.PostConstruct;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Component;

@Component
public class DatabaseConfigLogger {

    @Value("${spring.datasource.url}")
    private String datasourceUrl;

    @Value("${spring.datasource.username}")
    private String datasourceUsername;

    @Value("${spring.datasource.password}")
    private String datasourcePassword;

    @PostConstruct
    public void logDatabaseConnectionDetails() {
        System.out.println("Database URL: " + datasourceUrl);
        System.out.println("Database Username: " + datasourceUsername);
        System.out.println("Database Password: " + datasourcePassword);
    }
}