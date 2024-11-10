package br.com.josenaldo.wbu.services;

import br.com.josenaldo.wbu.entities.Account;
import br.com.josenaldo.wbu.entities.EntityId;
import br.com.josenaldo.wbu.exceptions.NotFoundException;
import br.com.josenaldo.wbu.repositories.AccountRepository;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class AccountService {

    private final AccountRepository accountRepository;

    public AccountService(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    public List<Account> findAll() {
        return accountRepository.findAll();
    }

    public Account findById(String accountId) {
        EntityId entityId = new EntityId(accountId);
        Optional<Account> account = accountRepository.findById(entityId);

        if(account.isEmpty()) {
            throw new NotFoundException("Account not found");
        }

        return account.get();
    }
}
