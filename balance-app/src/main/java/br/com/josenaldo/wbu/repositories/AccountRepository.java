package br.com.josenaldo.wbu.repositories;

import br.com.josenaldo.wbu.entities.Account;
import br.com.josenaldo.wbu.entities.EntityId;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AccountRepository extends JpaRepository<Account, EntityId> {
}