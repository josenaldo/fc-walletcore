package br.com.josenaldo.wbu;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;
import org.springframework.data.jpa.repository.config.EnableJpaRepositories;

@EnableJpaRepositories("br.com.josenaldo.wbu.repositories")
@EntityScan("br.com.josenaldo.wbu.entities")
@SpringBootApplication
public class FcWalletcoreBalanceUpdaterApplication {

	public static void main(String[] args) {
		SpringApplication.run(FcWalletcoreBalanceUpdaterApplication.class, args);
	}

}
